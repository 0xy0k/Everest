package keeper

import (
	appparams "github.com/TsukiCore/tsuki/app/params"
	"github.com/TsukiCore/tsuki/x/ubi/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper is for managing token module
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	bk       types.BankKeeper
	sk       types.SpendingKeeper
	dk       types.DistrKeeper
}

// NewKeeper returns instance of a keeper
func NewKeeper(storeKey storetypes.StoreKey, cdc codec.BinaryCodec, bk types.BankKeeper, sk types.SpendingKeeper, dk types.DistrKeeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		bk:       bk,
		sk:       sk,
		dk:       dk,
	}
}

// DefaultDenom returns the denom that is basically used for fee payment
func (k Keeper) DefaultDenom(ctx sdk.Context) string {
	return appparams.DefaultDenom
}
