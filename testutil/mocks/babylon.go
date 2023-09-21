// Code generated by MockGen. DO NOT EDIT.
// Source: clientcontroller/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/babylonchain/babylon/types"
	types0 "github.com/babylonchain/babylon/x/btcstaking/types"
	clientcontroller "github.com/babylonchain/btc-validator/clientcontroller"
	types1 "github.com/babylonchain/btc-validator/types"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	secp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	types2 "github.com/cosmos/cosmos-sdk/types"
	types3 "github.com/cosmos/cosmos-sdk/x/staking/types"
	provider "github.com/cosmos/relayer/v2/relayer/provider"
	gomock "github.com/golang/mock/gomock"
)

// MockClientController is a mock of ClientController interface.
type MockClientController struct {
	ctrl     *gomock.Controller
	recorder *MockClientControllerMockRecorder
}

// MockClientControllerMockRecorder is the mock recorder for MockClientController.
type MockClientControllerMockRecorder struct {
	mock *MockClientController
}

// NewMockClientController creates a new mock instance.
func NewMockClientController(ctrl *gomock.Controller) *MockClientController {
	mock := &MockClientController{ctrl: ctrl}
	mock.recorder = &MockClientControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientController) EXPECT() *MockClientControllerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClientController) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientControllerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientController)(nil).Close))
}

// CommitPubRandList mocks base method.
func (m *MockClientController) CommitPubRandList(btcPubKey *types.BIP340PubKey, startHeight uint64, pubRandList []types.SchnorrPubRand, sig *types.BIP340Signature) (*provider.RelayerTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitPubRandList", btcPubKey, startHeight, pubRandList, sig)
	ret0, _ := ret[0].(*provider.RelayerTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitPubRandList indicates an expected call of CommitPubRandList.
func (mr *MockClientControllerMockRecorder) CommitPubRandList(btcPubKey, startHeight, pubRandList, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitPubRandList", reflect.TypeOf((*MockClientController)(nil).CommitPubRandList), btcPubKey, startHeight, pubRandList, sig)
}

// GetStakingParams mocks base method.
func (m *MockClientController) GetStakingParams() (*clientcontroller.StakingParams, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStakingParams")
	ret0, _ := ret[0].(*clientcontroller.StakingParams)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStakingParams indicates an expected call of GetStakingParams.
func (mr *MockClientControllerMockRecorder) GetStakingParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStakingParams", reflect.TypeOf((*MockClientController)(nil).GetStakingParams))
}

// QueryBTCValidatorUnbondingDelegations mocks base method.
func (m *MockClientController) QueryBTCValidatorUnbondingDelegations(valBtcPk *types.BIP340PubKey, max uint64) ([]*types0.BTCDelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBTCValidatorUnbondingDelegations", valBtcPk, max)
	ret0, _ := ret[0].([]*types0.BTCDelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBTCValidatorUnbondingDelegations indicates an expected call of QueryBTCValidatorUnbondingDelegations.
func (mr *MockClientControllerMockRecorder) QueryBTCValidatorUnbondingDelegations(valBtcPk, max interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBTCValidatorUnbondingDelegations", reflect.TypeOf((*MockClientController)(nil).QueryBTCValidatorUnbondingDelegations), valBtcPk, max)
}

// QueryBestHeader mocks base method.
func (m *MockClientController) QueryBestHeader() (*coretypes.ResultHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBestHeader")
	ret0, _ := ret[0].(*coretypes.ResultHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBestHeader indicates an expected call of QueryBestHeader.
func (mr *MockClientControllerMockRecorder) QueryBestHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBestHeader", reflect.TypeOf((*MockClientController)(nil).QueryBestHeader))
}

// QueryBlockFinalization mocks base method.
func (m *MockClientController) QueryBlockFinalization(height uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBlockFinalization", height)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBlockFinalization indicates an expected call of QueryBlockFinalization.
func (mr *MockClientControllerMockRecorder) QueryBlockFinalization(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBlockFinalization", reflect.TypeOf((*MockClientController)(nil).QueryBlockFinalization), height)
}

// QueryBlocks mocks base method.
func (m *MockClientController) QueryBlocks(startHeight, endHeight, limit uint64) ([]*types1.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBlocks", startHeight, endHeight, limit)
	ret0, _ := ret[0].([]*types1.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBlocks indicates an expected call of QueryBlocks.
func (mr *MockClientControllerMockRecorder) QueryBlocks(startHeight, endHeight, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBlocks", reflect.TypeOf((*MockClientController)(nil).QueryBlocks), startHeight, endHeight, limit)
}

// QueryHeader mocks base method.
func (m *MockClientController) QueryHeader(height int64) (*coretypes.ResultHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryHeader", height)
	ret0, _ := ret[0].(*coretypes.ResultHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryHeader indicates an expected call of QueryHeader.
func (mr *MockClientControllerMockRecorder) QueryHeader(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryHeader", reflect.TypeOf((*MockClientController)(nil).QueryHeader), height)
}

// QueryHeightWithLastPubRand mocks base method.
func (m *MockClientController) QueryHeightWithLastPubRand(btcPubKey *types.BIP340PubKey) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryHeightWithLastPubRand", btcPubKey)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryHeightWithLastPubRand indicates an expected call of QueryHeightWithLastPubRand.
func (mr *MockClientControllerMockRecorder) QueryHeightWithLastPubRand(btcPubKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryHeightWithLastPubRand", reflect.TypeOf((*MockClientController)(nil).QueryHeightWithLastPubRand), btcPubKey)
}

// QueryLatestFinalizedBlocks mocks base method.
func (m *MockClientController) QueryLatestFinalizedBlocks(count uint64) ([]*types1.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLatestFinalizedBlocks", count)
	ret0, _ := ret[0].([]*types1.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLatestFinalizedBlocks indicates an expected call of QueryLatestFinalizedBlocks.
func (mr *MockClientControllerMockRecorder) QueryLatestFinalizedBlocks(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLatestFinalizedBlocks", reflect.TypeOf((*MockClientController)(nil).QueryLatestFinalizedBlocks), count)
}

// QueryNodeStatus mocks base method.
func (m *MockClientController) QueryNodeStatus() (*coretypes.ResultStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryNodeStatus")
	ret0, _ := ret[0].(*coretypes.ResultStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryNodeStatus indicates an expected call of QueryNodeStatus.
func (mr *MockClientControllerMockRecorder) QueryNodeStatus() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryNodeStatus", reflect.TypeOf((*MockClientController)(nil).QueryNodeStatus))
}

// QueryPendingBTCDelegations mocks base method.
func (m *MockClientController) QueryPendingBTCDelegations() ([]*types0.BTCDelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryPendingBTCDelegations")
	ret0, _ := ret[0].([]*types0.BTCDelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryPendingBTCDelegations indicates an expected call of QueryPendingBTCDelegations.
func (mr *MockClientControllerMockRecorder) QueryPendingBTCDelegations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryPendingBTCDelegations", reflect.TypeOf((*MockClientController)(nil).QueryPendingBTCDelegations))
}

// QueryUnbondindBTCDelegations mocks base method.
func (m *MockClientController) QueryUnbondindBTCDelegations() ([]*types0.BTCDelegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryUnbondindBTCDelegations")
	ret0, _ := ret[0].([]*types0.BTCDelegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryUnbondindBTCDelegations indicates an expected call of QueryUnbondindBTCDelegations.
func (mr *MockClientControllerMockRecorder) QueryUnbondindBTCDelegations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryUnbondindBTCDelegations", reflect.TypeOf((*MockClientController)(nil).QueryUnbondindBTCDelegations))
}

// QueryValidatorVotingPower mocks base method.
func (m *MockClientController) QueryValidatorVotingPower(btcPubKey *types.BIP340PubKey, blockHeight uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryValidatorVotingPower", btcPubKey, blockHeight)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryValidatorVotingPower indicates an expected call of QueryValidatorVotingPower.
func (mr *MockClientControllerMockRecorder) QueryValidatorVotingPower(btcPubKey, blockHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryValidatorVotingPower", reflect.TypeOf((*MockClientController)(nil).QueryValidatorVotingPower), btcPubKey, blockHeight)
}

// RegisterValidator mocks base method.
func (m *MockClientController) RegisterValidator(bbnPubKey *secp256k1.PubKey, btcPubKey *types.BIP340PubKey, pop *types0.ProofOfPossession, commission *types2.Dec, description *types3.Description) (*provider.RelayerTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterValidator", bbnPubKey, btcPubKey, pop, commission, description)
	ret0, _ := ret[0].(*provider.RelayerTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterValidator indicates an expected call of RegisterValidator.
func (mr *MockClientControllerMockRecorder) RegisterValidator(bbnPubKey, btcPubKey, pop, commission, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterValidator", reflect.TypeOf((*MockClientController)(nil).RegisterValidator), bbnPubKey, btcPubKey, pop, commission, description)
}

// SubmitBatchFinalitySigs mocks base method.
func (m *MockClientController) SubmitBatchFinalitySigs(btcPubKey *types.BIP340PubKey, blocks []*types1.BlockInfo, sigs []*types.SchnorrEOTSSig) (*provider.RelayerTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitBatchFinalitySigs", btcPubKey, blocks, sigs)
	ret0, _ := ret[0].(*provider.RelayerTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitBatchFinalitySigs indicates an expected call of SubmitBatchFinalitySigs.
func (mr *MockClientControllerMockRecorder) SubmitBatchFinalitySigs(btcPubKey, blocks, sigs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitBatchFinalitySigs", reflect.TypeOf((*MockClientController)(nil).SubmitBatchFinalitySigs), btcPubKey, blocks, sigs)
}

// SubmitFinalitySig mocks base method.
func (m *MockClientController) SubmitFinalitySig(btcPubKey *types.BIP340PubKey, blockHeight uint64, blockHash []byte, sig *types.SchnorrEOTSSig) (*provider.RelayerTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitFinalitySig", btcPubKey, blockHeight, blockHash, sig)
	ret0, _ := ret[0].(*provider.RelayerTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitFinalitySig indicates an expected call of SubmitFinalitySig.
func (mr *MockClientControllerMockRecorder) SubmitFinalitySig(btcPubKey, blockHeight, blockHash, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitFinalitySig", reflect.TypeOf((*MockClientController)(nil).SubmitFinalitySig), btcPubKey, blockHeight, blockHash, sig)
}

// SubmitJurySig mocks base method.
func (m *MockClientController) SubmitJurySig(btcPubKey, delPubKey *types.BIP340PubKey, stakingTxHash string, sig *types.BIP340Signature) (*provider.RelayerTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitJurySig", btcPubKey, delPubKey, stakingTxHash, sig)
	ret0, _ := ret[0].(*provider.RelayerTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitJurySig indicates an expected call of SubmitJurySig.
func (mr *MockClientControllerMockRecorder) SubmitJurySig(btcPubKey, delPubKey, stakingTxHash, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitJurySig", reflect.TypeOf((*MockClientController)(nil).SubmitJurySig), btcPubKey, delPubKey, stakingTxHash, sig)
}

// SubmitJuryUnbondingSigs mocks base method.
func (m *MockClientController) SubmitJuryUnbondingSigs(btcPubKey, delPubKey *types.BIP340PubKey, stakingTxHash string, unbondingSig, slashUnbondingSig *types.BIP340Signature) (*provider.RelayerTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitJuryUnbondingSigs", btcPubKey, delPubKey, stakingTxHash, unbondingSig, slashUnbondingSig)
	ret0, _ := ret[0].(*provider.RelayerTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitJuryUnbondingSigs indicates an expected call of SubmitJuryUnbondingSigs.
func (mr *MockClientControllerMockRecorder) SubmitJuryUnbondingSigs(btcPubKey, delPubKey, stakingTxHash, unbondingSig, slashUnbondingSig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitJuryUnbondingSigs", reflect.TypeOf((*MockClientController)(nil).SubmitJuryUnbondingSigs), btcPubKey, delPubKey, stakingTxHash, unbondingSig, slashUnbondingSig)
}

// SubmitValidatorUnbondingSig mocks base method.
func (m *MockClientController) SubmitValidatorUnbondingSig(valPubKey, delPubKey *types.BIP340PubKey, stakingTxHash string, sig *types.BIP340Signature) (*provider.RelayerTxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitValidatorUnbondingSig", valPubKey, delPubKey, stakingTxHash, sig)
	ret0, _ := ret[0].(*provider.RelayerTxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitValidatorUnbondingSig indicates an expected call of SubmitValidatorUnbondingSig.
func (mr *MockClientControllerMockRecorder) SubmitValidatorUnbondingSig(valPubKey, delPubKey, stakingTxHash, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitValidatorUnbondingSig", reflect.TypeOf((*MockClientController)(nil).SubmitValidatorUnbondingSig), valPubKey, delPubKey, stakingTxHash, sig)
}
