package main

import "github.com/cosmos/cosmos-sdk/crypto/keyring"

const (
	homeFlag           = "home"
	forceFlag          = "force"
	passphraseFlag     = "passphrase"
	fpPkFlag           = "finality-provider-pk"
	keyNameFlag        = "key-name"
	hdPathFlag         = "hd-path"
	chainIdFlag        = "chain-id"
	keyringBackendFlag = "keyring-backend"

	defaultKeyringBackend = keyring.BackendTest
	defaultHdPath         = ""
	defaultPassphrase     = ""
)