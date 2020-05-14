package createOrderBook

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Message{}, "tsukiHub/createOrderBook", nil)
}

var PackageCodec = codec.New()

func init() {
	RegisterCodec(PackageCodec)

}
