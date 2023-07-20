package types

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CustomGovKeeper defines the expected interface contract the tokens module requires
type CustomGovKeeper interface {
	CheckIfAllowedPermission(ctx sdk.Context, addr sdk.AccAddress, permValue govtypes.PermValue) bool
}
