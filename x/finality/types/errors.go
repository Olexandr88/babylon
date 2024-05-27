package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/finality module sentinel errors
var (
	ErrBlockNotFound         = errorsmod.Register(ModuleName, 1100, "Block is not found")
	ErrVoteNotFound          = errorsmod.Register(ModuleName, 1101, "vote is not found")
	ErrHeightTooHigh         = errorsmod.Register(ModuleName, 1102, "the chain has not reached the given height yet")
	ErrPubRandNotFound       = errorsmod.Register(ModuleName, 1103, "public randomness is not found")
	ErrPubRandCommitNotFound = errorsmod.Register(ModuleName, 1104, "public randomness commitment is not found")
	ErrNoPubRandYet          = errorsmod.Register(ModuleName, 1105, "the finality provider has not committed any public randomness yet")
	ErrTooFewPubRand         = errorsmod.Register(ModuleName, 1106, "the request contains too few public randomness")
	ErrInvalidPubRand        = errorsmod.Register(ModuleName, 1107, "the public randomness list is invalid")
	ErrEvidenceNotFound      = errorsmod.Register(ModuleName, 1108, "evidence is not found")
	ErrInvalidFinalitySig    = errorsmod.Register(ModuleName, 1109, "finality signature is not valid")
	ErrNoSlashableEvidence   = errorsmod.Register(ModuleName, 1110, "there is no slashable evidence")
)
