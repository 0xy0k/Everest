package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterRecoverySecret{}, "cosmos-sdk/MsgRegisterRecoverySecret", nil)
	cdc.RegisterConcrete(&MsgRotateRecoveryAddress{}, "cosmos-sdk/MsgRotateRecoveryAddress", nil)
	cdc.RegisterConcrete(&MsgIssueRecoveryTokens{}, "cosmos-sdk/MsgIssueRecoveryTokens", nil)
	cdc.RegisterConcrete(&MsgBurnRecoveryTokens{}, "cosmos-sdk/MsgBurnRecoveryTokens", nil)
	cdc.RegisterConcrete(&MsgClaimRRHolderRewards{}, "cosmos-sdk/MsgClaimRRHolderRewards", nil)
	cdc.RegisterConcrete(&MsgRegisterRRTokenHolder{}, "cosmos-sdk/MsgRegisterRRTokenHolder", nil)
}

// RegisterInterfaces register interfaces on registry
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterRecoverySecret{},
		&MsgRotateRecoveryAddress{},
		&MsgIssueRecoveryTokens{},
		&MsgBurnRecoveryTokens{},
		&MsgClaimRRHolderRewards{},
		&MsgRegisterRRTokenHolder{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/recovery module codec. Note, the codec
	// should ONLY be used in certain instances of tests and for JSON encoding as Amino
	// is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/recovery and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
