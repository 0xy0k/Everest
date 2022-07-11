package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec register codec and metadata
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgUpsertStakingPool{}, "tsukiHub/MsgUpsertStakingPool", nil)
	cdc.RegisterConcrete(&MsgDelegate{}, "tsukiHub/MsgDelegate", nil)
	cdc.RegisterConcrete(&MsgUndelegate{}, "tsukiHub/MsgUndelegate", nil)
	cdc.RegisterConcrete(&MsgClaimRewards{}, "tsukiHub/MsgClaimRewards", nil)
	cdc.RegisterConcrete(&MsgClaimUndelegation{}, "tsukiHub/MsgClaimUndelegation", nil)
	cdc.RegisterConcrete(&MsgRegisterDelegator{}, "tsukiHub/MsgRegisterDelegator", nil)
}

// RegisterInterfaces register Msg and structs
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpsertStakingPool{},
		&MsgDelegate{},
		&MsgUndelegate{},
		&MsgClaimRewards{},
		&MsgRegisterDelegator{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/staking module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/staking and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterCodec(amino)
	amino.Seal()
}
