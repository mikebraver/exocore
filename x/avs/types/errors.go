package types

// DONTCOVER

import (
	errorsmod "cosmossdk.io/errors"
)

// x/avs module sentinel errors
var (
	ErrNoKeyInTheStore = errorsmod.Register(ModuleName, 2, "there is no such key in the store")

	ErrAlreadyRegistered = errorsmod.Register(
		ModuleName, 3,
		"Error: Already registered",
	)
	ErrUnregisterNonExistent = errorsmod.Register(
		ModuleName, 4,
		"Error: No available avs to DeRegisterAction",
	)

	ErrInvalidAction = errorsmod.Register(
		ModuleName, 5,
		"Error: Undefined action",
	)
	ErrUnbondingPeriod = errorsmod.Register(
		ModuleName, 6,
		"Error: UnbondingPeriod check failed",
	)
	ErrEpochNotFound = errorsmod.Register(
		ModuleName, 7,
		"Error: epoch info not found",
	)
	ErrCallerAddressUnauthorized = errorsmod.Register(
		ModuleName, 8,
		"Error: caller address is not authorized",
	)
	ErrAvsNameMismatch = errorsmod.Register(
		ModuleName, 9,
		"Error: avs Name mismatch",
	)
	ErrNotYetRegistered = errorsmod.Register(
		ModuleName, 10,
		"Error: this AVS has not been registered yet",
	)
	ErrNotNull = errorsmod.Register(
		ModuleName, 11,
		"Error: this chainID shouldn't be null",
	)

	ErrInvalidAddr = errorsmod.Register(
		ModuleName, 12,
		"The address isn't a valid account address",
	)
	ErrAlreadyExists = errorsmod.Register(
		ModuleName, 13,
		"The task already exists",
	)
)
