package types

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
	"github.com/TsukiCore/cosmos-sdk/codec/types"
)

var (
	amino     = codec.New()
	ModuleCdc = codec.NewHybridCodec(amino, types.NewInterfaceRegistry())
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(&MsgClaimValidator{}, "tsukiHub/MsgClaimValidator", nil)
}
