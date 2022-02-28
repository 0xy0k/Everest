package types

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CustomGovKeeper interface {
	GetNetworkProperties(ctx sdk.Context) *govtypes.NetworkProperties
}
