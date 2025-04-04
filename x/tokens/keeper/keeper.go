package keeper

import (
	appparams "github.com/TsukiCore/tsuki/app/params"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// store prefixes
var (
	PrefixKeyTokenAlias      = []byte("token_alias_registry")
	PrefixKeyDenomToken      = []byte("denom_token_registry")
	PrefixKeyTokenRate       = []byte("token_rate_registry")
	PrefixKeyTokenBlackWhite = []byte("token_black_white")
)

// Keeper is for managing token module
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
}

// NewKeeper returns instance of a keeper
func NewKeeper(storeKey storetypes.StoreKey, cdc codec.BinaryCodec) Keeper {
	return Keeper{cdc: cdc, storeKey: storeKey}
}

// DefaultDenom returns the denom that is basically used for fee payment
func (k Keeper) DefaultDenom(ctx sdk.Context) string {
	return appparams.DefaultDenom
}
