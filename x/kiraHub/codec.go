package tsukiHub

import (
	"github.com/TsukiCore/cosmos-sdk/codec"

	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrderBook"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrder"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	createOrderBook.RegisterCodec(cdc)
	createOrder.RegisterCodec(cdc)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
