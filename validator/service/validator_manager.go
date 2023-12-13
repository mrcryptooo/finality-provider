package service

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/avast/retry-go/v4"
	bbntypes "github.com/babylonchain/babylon/types"
	btcstakingtypes "github.com/babylonchain/babylon/x/btcstaking/types"
	"go.uber.org/atomic"
	"go.uber.org/zap"

	"github.com/babylonchain/btc-validator/clientcontroller"
	"github.com/babylonchain/btc-validator/eotsmanager"
	"github.com/babylonchain/btc-validator/types"
	valcfg "github.com/babylonchain/btc-validator/validator/config"
	"github.com/babylonchain/btc-validator/validator/proto"
	valstore "github.com/babylonchain/btc-validator/validator/store"
)

const instanceTerminatingMsg = "terminating the validator instance due to critical error"

type CriticalError struct {
	err      error
	valBtcPk *bbntypes.BIP340PubKey
}

func (ce *CriticalError) Error() string {
	return fmt.Sprintf("critical err on validator %s: %s", ce.valBtcPk.MarshalHex(), ce.err.Error())
}

type ValidatorManager struct {
	isStarted *atomic.Bool

	mu sync.Mutex
	wg sync.WaitGroup

	// running validator instances map keyed by the hex string of the BTC public key
	vals map[string]*ValidatorInstance

	// needed for initiating validator instances
	vs     *valstore.ValidatorStore
	config *valcfg.Config
	cc     clientcontroller.ClientController
	em     eotsmanager.EOTSManager
	logger *zap.Logger

	criticalErrChan chan *CriticalError

	quit chan struct{}
}

func NewValidatorManager(vs *valstore.ValidatorStore,
	config *valcfg.Config,
	cc clientcontroller.ClientController,
	em eotsmanager.EOTSManager,
	logger *zap.Logger,
) (*ValidatorManager, error) {
	return &ValidatorManager{
		vals:            make(map[string]*ValidatorInstance),
		criticalErrChan: make(chan *CriticalError),
		isStarted:       atomic.NewBool(false),
		vs:              vs,
		config:          config,
		cc:              cc,
		em:              em,
		logger:          logger,
		quit:            make(chan struct{}),
	}, nil
}

// monitorCriticalErr takes actions when it receives critical errors from a validator instance
// if the validator is slashed, it will be terminated and the program keeps running in case
// new validators join
// otherwise, the program will panic
func (vm *ValidatorManager) monitorCriticalErr() {
	defer vm.wg.Done()

	var criticalErr *CriticalError
	for {
		select {
		case criticalErr = <-vm.criticalErrChan:
			vi, err := vm.GetValidatorInstance(criticalErr.valBtcPk)
			if err != nil {
				panic(fmt.Errorf("failed to get the validator instance: %w", err))
			}
			if errors.Is(criticalErr.err, btcstakingtypes.ErrBTCValAlreadySlashed) {
				vm.setValidatorSlashed(vi)
				vm.logger.Debug("the validator has been slashed",
					zap.String("pk", criticalErr.valBtcPk.MarshalHex()))
				continue
			}
			vi.logger.Fatal(instanceTerminatingMsg,
				zap.String("pk", criticalErr.valBtcPk.MarshalHex()), zap.Error(criticalErr.err))
		case <-vm.quit:
			return
		}
	}
}

// monitorStatusUpdate periodically check the status of each managed validators and update
// it accordingly. We update the status by querying the latest voting power and the slashed_height.
// In particular, we perform the following status transitions (REGISTERED, ACTIVE, INACTIVE, SLASHED):
// 1. if power == 0 and slashed_height == 0, if status == ACTIVE, change to INACTIVE, otherwise remain the same
// 2. if power == 0 and slashed_height > 0, set status to SLASHED and stop and remove the validator instance
// 3. if power > 0 (slashed_height must > 0), set status to ACTIVE
// NOTE: once error occurs, we log and continue as the status update is not critical to the entire program
func (vm *ValidatorManager) monitorStatusUpdate() {
	defer vm.wg.Done()

	if vm.config.StatusUpdateInterval == 0 {
		vm.logger.Info("the status update is disabled")
		return
	}

	statusUpdateTicker := time.NewTicker(vm.config.StatusUpdateInterval)
	defer statusUpdateTicker.Stop()

	for {
		select {
		case <-statusUpdateTicker.C:
			latestBlock, err := vm.getLatestBlockWithRetry()
			if err != nil {
				vm.logger.Debug("failed to get the latest block", zap.Error(err))
				continue
			}
			vals := vm.ListValidatorInstances()
			for _, v := range vals {
				oldStatus := v.GetStatus()
				power, err := v.GetVotingPowerWithRetry(latestBlock.Height)
				if err != nil {
					vm.logger.Debug(
						"failed to get the voting power",
						zap.String("val_btc_pk", v.GetBtcPkHex()),
						zap.Uint64("height", latestBlock.Height),
						zap.Error(err),
					)
					continue
				}
				// power > 0 (slashed_height must > 0), set status to ACTIVE
				if power > 0 {
					if oldStatus != proto.ValidatorStatus_ACTIVE {
						v.MustSetStatus(proto.ValidatorStatus_ACTIVE)
						vm.logger.Debug(
							"the validator status is changed to ACTIVE",
							zap.String("val_btc_pk", v.GetBtcPkHex()),
							zap.String("old_status", oldStatus.String()),
							zap.Uint64("power", power),
						)
					}
					continue
				}
				slashed, err := v.GetValidatorSlashedWithRetry()
				if err != nil {
					vm.logger.Debug(
						"failed to get the slashed height",
						zap.String("val_btc_pk", v.GetBtcPkHex()),
						zap.Error(err),
					)
					continue
				}
				// power == 0 and slashed == true, set status to SLASHED and stop and remove the validator instance
				if slashed {
					vm.setValidatorSlashed(v)
					vm.logger.Debug(
						"the validator is slashed",
						zap.String("val_btc_pk", v.GetBtcPkHex()),
						zap.String("old_status", oldStatus.String()),
					)
					continue
				}
				// power == 0 and slashed_height == 0, change to INACTIVE if the current status is ACTIVE
				if oldStatus == proto.ValidatorStatus_ACTIVE {
					v.MustSetStatus(proto.ValidatorStatus_INACTIVE)
					vm.logger.Debug(
						"the validator status is changed to INACTIVE",
						zap.String("val_btc_pk", v.GetBtcPkHex()),
						zap.String("old_status", oldStatus.String()),
					)
				}
			}
		case <-vm.quit:
			return
		}
	}
}

func (vm *ValidatorManager) setValidatorSlashed(vi *ValidatorInstance) {
	vi.MustSetStatus(proto.ValidatorStatus_SLASHED)
	if err := vm.removeValidatorInstance(vi.GetBtcPkBIP340()); err != nil {
		panic(fmt.Errorf("failed to terminate a slashed validator %s: %w", vi.GetBtcPkHex(), err))
	}
}

func (vm *ValidatorManager) StartValidator(valPk *bbntypes.BIP340PubKey, passphrase string) error {
	if !vm.isStarted.Load() {
		vm.isStarted.Store(true)

		vm.wg.Add(1)
		go vm.monitorCriticalErr()

		vm.wg.Add(1)
		go vm.monitorStatusUpdate()
	}

	if vm.numOfRunningValidators() >= int(vm.config.MaxNumValidators) {
		return fmt.Errorf("reaching maximum number of running validators %v", vm.config.MaxNumValidators)
	}

	if err := vm.addValidatorInstance(valPk, passphrase); err != nil {
		return err
	}

	return nil
}

func (vm *ValidatorManager) Stop() error {
	if !vm.isStarted.Swap(false) {
		return fmt.Errorf("the validator manager has already stopped")
	}

	var stopErr error

	for _, v := range vm.vals {
		if !v.IsRunning() {
			continue
		}
		if err := v.Stop(); err != nil {
			stopErr = err
			break
		}
	}

	close(vm.quit)
	vm.wg.Wait()

	return stopErr
}

func (vm *ValidatorManager) ListValidatorInstances() []*ValidatorInstance {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	valsList := make([]*ValidatorInstance, 0, len(vm.vals))
	for _, v := range vm.vals {
		valsList = append(valsList, v)
	}

	return valsList
}

func (vm *ValidatorManager) GetValidatorInstance(valPk *bbntypes.BIP340PubKey) (*ValidatorInstance, error) {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	keyHex := valPk.MarshalHex()
	v, exists := vm.vals[keyHex]
	if !exists {
		return nil, fmt.Errorf("cannot find the validator instance with PK: %s", keyHex)
	}

	return v, nil
}

func (vm *ValidatorManager) removeValidatorInstance(valPk *bbntypes.BIP340PubKey) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	keyHex := valPk.MarshalHex()
	v, exists := vm.vals[keyHex]
	if !exists {
		return fmt.Errorf("cannot find the validator instance with PK: %s", keyHex)
	}
	if v.IsRunning() {
		if err := v.Stop(); err != nil {
			return fmt.Errorf("failed to stop the validator instance %s", keyHex)
		}
	}

	delete(vm.vals, keyHex)
	return nil
}

func (vm *ValidatorManager) numOfRunningValidators() int {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	return len(vm.vals)
}

// addValidatorInstance creates a validator instance, starts it and adds it into the validator manager
func (vm *ValidatorManager) addValidatorInstance(
	pk *bbntypes.BIP340PubKey,
	passphrase string,
) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	pkHex := pk.MarshalHex()
	if _, exists := vm.vals[pkHex]; exists {
		return fmt.Errorf("validator instance already exists")
	}

	valIns, err := NewValidatorInstance(pk, vm.config, vm.vs, vm.cc, vm.em, passphrase, vm.criticalErrChan, vm.logger)
	if err != nil {
		return fmt.Errorf("failed to create validator %s instance: %w", pkHex, err)
	}

	if err := valIns.Start(); err != nil {
		return fmt.Errorf("failed to start validator %s instance: %w", pkHex, err)
	}

	vm.vals[pkHex] = valIns

	return nil
}

func (vm *ValidatorManager) getLatestBlockWithRetry() (*types.BlockInfo, error) {
	var (
		latestBlock *types.BlockInfo
		err         error
	)

	if err := retry.Do(func() error {
		latestBlock, err = vm.cc.QueryBestBlock()
		if err != nil {
			return err
		}
		return nil
	}, RtyAtt, RtyDel, RtyErr, retry.OnRetry(func(n uint, err error) {
		vm.logger.Debug(
			"failed to query the consumer chain for the latest block",
			zap.Uint("attempt", n+1),
			zap.Uint("max_attempts", RtyAttNum),
			zap.Error(err),
		)
	})); err != nil {
		return nil, err
	}

	return latestBlock, nil
}