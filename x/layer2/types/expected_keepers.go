package types

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	spendingtypes "github.com/TsukiCore/tsuki/x/spending/types"
	stakingtypes "github.com/TsukiCore/tsuki/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CustomGovKeeper interface {
	GetNetworkProperties(ctx sdk.Context) *govtypes.NetworkProperties
}

type BankKeeper interface {
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}

type StakingKeeper interface {
	GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (stakingtypes.Validator, error)
	GetValidator(sdk.Context, sdk.ValAddress) (stakingtypes.Validator, error)
}

type SpendingKeeper interface {
	CreateSpendingPool(ctx sdk.Context, pool spendingtypes.SpendingPool) error
	DepositSpendingPoolFromModule(ctx sdk.Context, moduleName, poolName string, amount sdk.Coins) error
}
