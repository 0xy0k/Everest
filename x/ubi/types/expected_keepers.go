package types

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	spendingtypes "github.com/TsukiCore/tsuki/x/spending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CustomGovKeeper interface {
	GetNetworkProperties(ctx sdk.Context) *govtypes.NetworkProperties
}

type BankKeeper interface {
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
}

type SpendingKeeper interface {
	GetSpendingPool(ctx sdk.Context, name string) *spendingtypes.SpendingPool
	DepositSpendingPoolFromModule(ctx sdk.Context, moduleName, poolName string, amounts sdk.Coins) error
}

type DistrKeeper interface {
	InflationPossible(ctx sdk.Context) bool
}
