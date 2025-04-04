package keeper

import (
	"github.com/TsukiCore/tsuki/x/basket/types"
	govkeeper "github.com/TsukiCore/tsuki/x/gov/keeper"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	tokenskeeper "github.com/TsukiCore/tsuki/x/tokens/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper is for managing token module
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey
	ak       types.AccountKeeper
	bk       types.BankKeeper
	gk       govkeeper.Keeper
	tk       tokenskeeper.Keeper
	mk       types.MultiStakingKeeper
}

// NewKeeper returns instance of a keeper
func NewKeeper(storeKey storetypes.StoreKey, cdc codec.BinaryCodec, ak types.AccountKeeper, bk types.BankKeeper, gk govkeeper.Keeper, tk tokenskeeper.Keeper, mk types.MultiStakingKeeper) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		ak:       ak,
		bk:       bk,
		gk:       gk,
		tk:       tk,
		mk:       mk,
	}
}

func (k Keeper) CheckIfAllowedPermission(ctx sdk.Context, addr sdk.AccAddress, permValue govtypes.PermValue) bool {
	return govkeeper.CheckIfAllowedPermission(ctx, k.gk, addr, govtypes.PermHandleBasketEmergency)
}
