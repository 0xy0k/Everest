package types

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	multistakingtypes "github.com/TsukiCore/tsuki/x/multistaking/types"
	spendingtypes "github.com/TsukiCore/tsuki/x/spending/types"
	stakingtypes "github.com/TsukiCore/tsuki/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type CustomGovKeeper interface {
	GetNetworkProperties(ctx sdk.Context) *govtypes.NetworkProperties
}

type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
}

type AccountKeeper interface {
	GetModuleAccount(ctx sdk.Context, moduleName string) authtypes.ModuleAccountI
}

type SpendingKeeper interface {
	GetSpendingPool(ctx sdk.Context, name string) *spendingtypes.SpendingPool
	DepositSpendingPoolFromModule(ctx sdk.Context, moduleName, poolName string, amount sdk.Coin) error
}

type StakingKeeper interface {
	GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (stakingtypes.Validator, error)
	GetValidator(sdk.Context, sdk.ValAddress) (stakingtypes.Validator, error)
}

type MultiStakingKeeper interface {
	GetStakingPoolByValidator(ctx sdk.Context, validator string) (pool multistakingtypes.StakingPool, found bool)
	IncreasePoolRewards(ctx sdk.Context, pool multistakingtypes.StakingPool, rewards sdk.Coins)
}
