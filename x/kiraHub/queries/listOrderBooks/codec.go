package listOrderBooks

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
)


func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(QueryListOrderBooks{}, "testtoken/query", nil)
}

var packageCodec = codec.New()

func init() {
	RegisterCodec(packageCodec)
	packageCodec.Seal()
}