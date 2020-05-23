package types

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(OrderBook{}, "tsukiHub/OrderBook", nil)
	cdc.RegisterConcrete(LimitOrder{}, "tsukiHub/LimitOrder", nil)
}

var PackageCodec = codec.New()

func init() {
	RegisterCodec(PackageCodec)
}
