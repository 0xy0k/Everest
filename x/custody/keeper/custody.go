package keeper

import (
	"github.com/TsukiCore/tsuki/x/custody/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetCustodyInfoByAddress(ctx sdk.Context, address sdk.AccAddress) *types.CustodySettings {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PrefixKeyCustodyRecord))
	bz := prefixStore.Get(address)

	if bz == nil {
		return nil
	}

	info := new(types.CustodySettings)
	k.cdc.MustUnmarshal(bz, info)

	return info
}

func (k Keeper) SetCustodyRecord(ctx sdk.Context, record types.CustodyRecord) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyCustodyRecord), record.Address...)

	store.Set(key, k.cdc.MustMarshal(&record.CustodySettings))
}

func (k Keeper) GetCustodyCustodiansByAddress(ctx sdk.Context, address sdk.AccAddress) *types.CustodyCustodianList {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PrefixKeyCustodyCustodians))
	bz := prefixStore.Get(address)

	if bz == nil {
		return nil
	}

	info := new(types.CustodyCustodianList)
	k.cdc.MustUnmarshal(bz, info)

	return info
}

func (k Keeper) AddToCustodyCustodians(ctx sdk.Context, record types.CustodyCustodiansRecord) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyCustodyCustodians), record.Address...)

	store.Set(key, k.cdc.MustMarshal(record.CustodyCustodians))
}

func (k Keeper) DropCustodyCustodiansByAddress(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyCustodyCustodians), address...)

	store.Delete(key)
}

func (k Keeper) GetCustodyWhiteListByAddress(ctx sdk.Context, address sdk.AccAddress) *types.CustodyWhiteList {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PrefixKeyCustodyWhiteList))
	bz := prefixStore.Get(address)

	if bz == nil {
		return nil
	}

	info := new(types.CustodyWhiteList)
	k.cdc.MustUnmarshal(bz, info)

	return info
}

func (k Keeper) AddToCustodyWhiteList(ctx sdk.Context, record types.CustodyWhiteListRecord) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyCustodyWhiteList), record.Address...)

	store.Set(key, k.cdc.MustMarshal(record.CustodyWhiteList))
}

func (k Keeper) DropCustodyWhiteListByAddress(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyCustodyWhiteList), address...)

	store.Delete(key)
}

func (k Keeper) GetCustodyLimitsByAddress(ctx sdk.Context, address sdk.AccAddress) *types.CustodyLimits {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PrefixKeyCustodyLimits))
	bz := prefixStore.Get(address)

	if bz == nil {
		return nil
	}

	info := new(types.CustodyLimits)
	k.cdc.MustUnmarshal(bz, info)

	return info
}

func (k Keeper) AddToCustodyLimits(ctx sdk.Context, record types.CustodyLimitRecord) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyCustodyLimits), record.Address...)

	store.Set(key, k.cdc.MustMarshal(record.CustodyLimits))
}

func (k Keeper) DropCustodyLimitsByAddress(ctx sdk.Context, address sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyCustodyLimits), address...)

	store.Delete(key)
}

func (k Keeper) GetCustodyLimitsStatusByAddress(ctx sdk.Context, address sdk.AccAddress) *types.CustodyStatuses {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PrefixKeyCustodyLimitsStatus))
	bz := prefixStore.Get(address)

	if bz == nil {
		return nil
	}

	info := new(types.CustodyStatuses)
	k.cdc.MustUnmarshal(bz, info)

	return info
}

func (k Keeper) AddToCustodyLimitsStatus(ctx sdk.Context, record types.CustodyLimitStatusRecord) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyCustodyLimitsStatus), record.Address...)

	store.Set(key, k.cdc.MustMarshal(record.CustodyStatuses))
}
