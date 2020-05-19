package tsukiHub

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrderBook"
)

type Keeper interface {
	getCreateOrderBookKeeper() createOrderBook.Keeper
}

type baseKeeper struct {
	createOrderBookKeeper   createOrderBook.Keeper
}

func NewKeeper(codec *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return baseKeeper{
		createOrderBookKeeper:   createOrderBook.NewKeeper(codec, storeKey),
	}
}

func (baseKeeper baseKeeper) getCreateOrderBookKeeper() createOrderBook.Keeper { return baseKeeper.createOrderBookKeeper }






