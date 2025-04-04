package types

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	multistakingtypes "github.com/TsukiCore/tsuki/x/multistaking/types"
	stakingtypes "github.com/TsukiCore/tsuki/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// ParamSubspace defines the expected Subspace interfacace
type ParamSubspace interface {
	HasKeyTable() bool
	WithKeyTable(table paramtypes.KeyTable) paramtypes.Subspace
	Get(ctx sdk.Context, key []byte, ptr interface{})
	GetParamSet(ctx sdk.Context, ps paramtypes.ParamSet)
	SetParamSet(ctx sdk.Context, ps paramtypes.ParamSet)
}

// StakingKeeper expected staking keeper
type StakingKeeper interface {
	// iterate through validators by operator address, execute func for each validator
	IterateValidators(sdk.Context,
		func(index int64, validator *stakingtypes.Validator) (stop bool))

	GetValidator(sdk.Context, sdk.ValAddress) (stakingtypes.Validator, error)            // get a particular validator by operator address
	GetValidatorByConsAddr(sdk.Context, sdk.ConsAddress) (stakingtypes.Validator, error) // get a particular validator by consensus address
	GetValidatorSet(ctx sdk.Context) []stakingtypes.Validator                            // get all validator set

	// activate/inactivate the validator and delegators of the validator, specifying offence height, offence power, and slash fraction
	Inactivate(sdk.Context, sdk.ValAddress) error // inactivate a validator
	Activate(sdk.Context, sdk.ValAddress) error   // activate a validator
	Jail(sdk.Context, sdk.ValAddress) error       // jail a validator
	ResetWholeValidatorRank(sdk.Context)          // reset whole validator rank

	// pause/unpause the validator and delegators of the validator, specifying offence height, offence power, and slash fraction
	Pause(sdk.Context, sdk.ValAddress) error   // pause a validator
	Unpause(sdk.Context, sdk.ValAddress) error // unpause a validator

	HandleValidatorSignature(sdk.Context, sdk.ValAddress, bool, int64) error

	// MaxValidators returns the maximum amount of joined validators
	MaxValidators(sdk.Context) uint32

	GetIdRecordsByAddress(sdk.Context, sdk.AccAddress) []govtypes.IdentityRecord
}

// SlashingHooks event hooks for slashing
type SlashingHooks interface {
	AfterSlashProposalRaise(ctx sdk.Context, valAddr sdk.ValAddress, pool multistakingtypes.StakingPool)
}

// StakingHooks event hooks for staking validator object (noalias)
type StakingHooks interface {
	AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress)                           // Must be called when a validator is created
	AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) // Must be called when a validator is deleted
	AfterValidatorJoined(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress)  // Must be called when a validator is joined
}

// GovKeeper expected governance keeper
type GovKeeper interface {
	GetNetworkProperties(sdk.Context) *govtypes.NetworkProperties // returns network properties
	GetProposals(ctx sdk.Context) ([]govtypes.Proposal, error)
	SaveProposal(ctx sdk.Context, proposal govtypes.Proposal)
	CreateAndSaveProposalWithContent(ctx sdk.Context, title, description string, content govtypes.Content) (uint64, error)
}

type MultiStakingKeeper interface {
	GetStakingPoolByValidator(ctx sdk.Context, validator string) (pool multistakingtypes.StakingPool, found bool)
	GetAllStakingPools(ctx sdk.Context) []multistakingtypes.StakingPool
	SlashStakingPool(ctx sdk.Context, validator string, slash sdk.Dec)
}
