package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*Content)(nil), nil)

	registerPermissionsCodec(cdc)
	registerRolesCodec(cdc)
	registerCouncilorCodec(cdc)
	registerProposalCodec(cdc)

	cdc.RegisterConcrete(&MsgSetNetworkProperties{}, "tsukiHub/MsgSetNetworkProperties", nil)
	cdc.RegisterConcrete(&MsgSetExecutionFee{}, "tsukiHub/MsgSetExecutionFee", nil)
}

func registerProposalCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgProposalAssignPermission{}, "tsukiHub/MsgProposalAssignPermission", nil)
	cdc.RegisterConcrete(&MsgProposalSetNetworkProperty{}, "tsukiHub/MsgProposalSetNetworkProperty", nil)
	cdc.RegisterConcrete(&MsgProposalUpsertDataRegistry{}, "tsukiHub/MsgProposalUpsertDataRegistry", nil)
	cdc.RegisterConcrete(&MsgVoteProposal{}, "tsukiHub/MsgVoteProposal", nil)
}

func registerCouncilorCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgClaimCouncilor{}, "tsukiHub/MsgClaimCouncilor", nil)
}

func registerPermissionsCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgWhitelistPermissions{}, "tsukiHub/MsgWhitelistPermissions", nil)
	cdc.RegisterConcrete(&MsgBlacklistPermissions{}, "tsukiHub/MsgBlacklistPermissions", nil)
}

func registerRolesCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateRole{}, "tsukiHub/MsgCreateRole", nil)
	cdc.RegisterConcrete(&MsgAssignRole{}, "tsukiHub/MsgAssignRole", nil)
	cdc.RegisterConcrete(&MsgRemoveRole{}, "tsukiHub/MsgRemoveRole", nil)

	cdc.RegisterConcrete(&MsgWhitelistRolePermission{}, "tsukiHub/MsgWhitelistRolePermission", nil)
	cdc.RegisterConcrete(&MsgBlacklistRolePermission{}, "tsukiHub/MsgBlacklistRolePermission", nil)
	cdc.RegisterConcrete(&MsgRemoveWhitelistRolePermission{}, "tsukiHub/MsgRemoveWhitelistRolePermission", nil)
	cdc.RegisterConcrete(&MsgRemoveBlacklistRolePermission{}, "tsukiHub/MsgRemoveBlacklistRolePermission", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWhitelistPermissions{},
		&MsgBlacklistPermissions{},

		&MsgSetNetworkProperties{},
		&MsgSetExecutionFee{},

		&MsgClaimCouncilor{},

		&MsgAssignRole{},
		&MsgCreateRole{},
		&MsgRemoveRole{},

		&MsgWhitelistRolePermission{},
		&MsgBlacklistRolePermission{},
		&MsgRemoveWhitelistRolePermission{},
		&MsgRemoveBlacklistRolePermission{},

		&MsgProposalAssignPermission{},
		&MsgProposalSetNetworkProperty{},
		&MsgProposalUpsertDataRegistry{},
		&MsgVoteProposal{},
	)

	registry.RegisterInterface(
		"tsuki.gov.Content",
		(*Content)(nil),
		&AssignPermissionProposal{},
		&SetNetworkPropertyProposal{},
		&UpsertDataRegistryProposal{},
	)

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
