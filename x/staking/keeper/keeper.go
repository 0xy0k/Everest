package keeper

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	types2 "github.com/TsukiCore/tsuki/x/staking/types"
)

// Keeper represents the keeper that maintains the Validator Registry.
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.Marshaler
}

// NewKeeper returns new keeper.
func NewKeeper(storeKey sdk.StoreKey, cdc codec.Marshaler) *Keeper {
	return &Keeper{storeKey: storeKey, cdc: cdc}
}

func (k Keeper) AddValidator(ctx sdk.Context, validator types2.Validator) {
}
