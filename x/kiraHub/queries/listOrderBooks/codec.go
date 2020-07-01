package listOrderBooks

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
)


func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(QueryListOrderBooks{}, "tsukiHub/queryOrderBooks", nil)
	cdc.RegisterConcrete(QueryListOrderBooksByTP{}, "tsukiHub/queryOrderBooksByTP", nil)
}

var packageCodec = codec.New()

func init() {
	RegisterCodec(packageCodec)
	packageCodec.Seal()
}