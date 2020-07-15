package createOrder

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
)

func RegisterCodec(codec *codec.Codec) {
	codec.RegisterConcrete(Message{}, "tsukiHub/createOrder", nil)
}

var PackageCodec = codec.New()

func init() {
	RegisterCodec(PackageCodec)
}