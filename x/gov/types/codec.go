package types

import (
	functionmeta "github.com/TsukiCore/tsuki/function_meta"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterCodec register codec and metadata
func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*Content)(nil), nil)

	registerPermissionsCodec(cdc)
	registerRolesCodec(cdc)
	registerCouncilorCodec(cdc)
	registerProposalCodec(cdc)
	registerIdRecordsCodec(cdc)

	cdc.RegisterConcrete(&MsgSetNetworkProperties{}, "tsukiHub/MsgSetNetworkProperties", nil)
	functionmeta.AddNewFunction((&MsgSetNetworkProperties{}).Type(), `{
		"description": "MsgSetNetworkProperties defines a message to set network properties with specific permission.",
		"parameters": {
			"network_properties": {
				"type":        "<Object>",
				"description": "network properties to be set.",
				"fields": {
					"min_tx_fee": {
						"type":        "uint64",
						"description": "minimum transaction fee"
					},
					"max_tx_fee": {
						"type":        "uint64",
						"description": "maximum transaction fee"
					},
					"vote_quorum": {
						"type":        "uint64",
						"description": "vote quorum"
					},
					"proposal_end_time": {
						"type":        "uint64",
						"description": "proposal end time"
					},
					"proposal_enactment_time": {
						"type":        "uint64",
						"description": "proposal enactment time"
					},
					"enable_foreign_fee_payments": {
						"type":        "bool",
						"description": "flag to show if foreign fee payment is enabled"
					}
				}
			},
			"proposer": {
				"type":        "address",
				"description": "proposer who propose this message."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgSetExecutionFee{}, "tsukiHub/MsgSetExecutionFee", nil)
	functionmeta.AddNewFunction((&MsgSetExecutionFee{}).Type(), `{
		"description": "MsgSetExecutionFee defines a message to set execution fee with specific permission.",
		"parameters": {
			"transaction_type": {
				"type":        "string",
				"description": "Type of the transaction that given permission allows to execute"
			},
			"execution_fee": {
				"type":        "uint64",
				"description": "How much user should pay for executing this specific function"
			},
			"failure_fee": {
				"type":        "uint64",
				"description": "How much user should pay if function fails to execute"
			},
			"timeout": {
				"type":        "uint64",
				"description": "After what time function execution should fail"
			},
			"default_parameters": {
				"type":        "bool",
				"description": "Default values that the function in question will consume as input parameters before execution"
			},
			"proposer": {
				"type":        "address",
				"description": "proposer who propose this message."
			}
		}
	}`)
}

func registerProposalCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgSubmitProposal{}, "tsukiHub/MsgSubmitProposal", nil)
	functionmeta.AddNewFunction((&MsgSubmitProposal{}).Type(), `{
		"description": "MsgSubmitProposal defines a proposal message to submit a proposal.",
		"parameters": {
			"proposer": {
				"type":        "address",
				"description": "the proposer of the proposal."
			},
			"title": {
				"type":        "string",
				"description": "the title of the proposal."
			},
			"description": {
				"type":        "string",
				"description": "the description of the proposal."
			},
			"content": {
				"type":        "object",
				"description": "the content of the proposal - different by type of proposal"
			}
		}
	}`)

	cdc.RegisterConcrete(&MsgVoteProposal{}, "tsukiHub/MsgVoteProposal", nil)
	functionmeta.AddNewFunction((&MsgVoteProposal{}).Type(), `{
		"description": "MsgVoteProposal defines a proposal message to vote on a submitted proposal.",
		"parameters": {
			"proposal_id": {
				"type":        "uint64",
				"description": "id of proposal to be voted."
			},
			"voter": {
				"type":        "address",
				"description": "the address of the voter who vote on the proposal."
			},
			"value": {
				"type":        "enum<VoteOption>",
				"description": "vote option: [yes, no, veto, abstain]"
			}
		}
	}`)
}

func registerIdRecordsCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterIdentityRecords{}, "tsukiHub/MsgRegisterIdentityRecords", nil)
	functionmeta.AddNewFunction((&MsgRegisterIdentityRecords{}).Type(), `{
		"description": "MsgRegisterIdentityRecords defines a proposal message to create a identity record.",
		"parameters": {
			"address": {
				"type":        "string",
				"description": "the address for the identity record."
			},
			"infos": {
				"type":        "array",
				"description": "key/value array for the mappings of the identity record."
			}
		}
	}`)

	cdc.RegisterConcrete(&MsgDeleteIdentityRecords{}, "tsukiHub/MsgDeleteIdentityRecords", nil)
	functionmeta.AddNewFunction((&MsgDeleteIdentityRecords{}).Type(), `{
		"description": "MsgDeleteIdentityRecords defines a method to delete identity records owned by an address.",
		"parameters": {
			"address": {
				"type":        "string",
				"description": "the address of requester."
			},
			"keys": {
				"type":        "array",
				"description": "the array string that defines identity record key values to be deleted."
			}
		}
	}`)

	cdc.RegisterConcrete(&MsgRequestIdentityRecordsVerify{}, "tsukiHub/MsgRequestIdentityRecordsVerify", nil)
	functionmeta.AddNewFunction((&MsgRequestIdentityRecordsVerify{}).Type(), `{
		"description": "MsgRequestIdentityRecordsVerify defines a proposal message to request an identity record verification from a specific verifier.",
		"parameters": {
			"address": {
				"type":        "string",
				"description": "the address of requester."
			},
			"verifier": {
				"type":        "string",
				"description": "the address of verifier."
			},
			"record_ids": {
				"type":        "array<uint64>",
				"description": "the id of records to be verified."
			},
			"tip": {
				"type":        "coins",
				"description": "the amount of coins to be given up-on accepting the request."
			}
		}
	}`)

	cdc.RegisterConcrete(&MsgHandleIdentityRecordsVerifyRequest{}, "tsukiHub/MsgHandleIdentityRecordsVerifyRequest", nil)
	functionmeta.AddNewFunction((&MsgHandleIdentityRecordsVerifyRequest{}).Type(), `{
		"description": "MsgHandleIdentityRecordsVerifyRequest defines a proposal message to approve or reject an identity record request.",
		"parameters": {
			"verifier": {
				"type":        "string",
				"description": "the address of verifier."
			},
			"verify_request_id": {
				"type":        "uint64",
				"description": "the id of verification request."
			},
			"yes": {
				"type": "bool",
				"optional": true,
				"description": "defines approval or rejecting an identity request (default false)"
			}
		}
	}`)

	cdc.RegisterConcrete(&MsgCancelIdentityRecordsVerifyRequest{}, "tsukiHub/MsgCancelIdentityRecordsVerifyRequest", nil)
	functionmeta.AddNewFunction((&MsgCancelIdentityRecordsVerifyRequest{}).Type(), `{
		"description": "MsgCancelIdentityRecordsVerifyRequest defines a proposal message to cancel an identity record request.",
		"parameters": {
			"executor": {
				"type":        "string",
				"description": "the address of requester."
			},
			"verify_request_id": {
				"type":        "uint64",
				"description": "the id of verification request."
			}
		}
	}`)
}

func registerCouncilorCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgClaimCouncilor{}, "tsukiHub/MsgClaimCouncilor", nil)
	functionmeta.AddNewFunction((&MsgClaimCouncilor{}).Type(), `{
		"description": "MsgClaimCouncilor defines a message to claim councilor when the proposer.",
		"parameters": {
			"moniker": {
				"type":        "string",
				"description": "validator's name or nickname."
			},
			"address": {
				"type":        "string",
				"description": "Address to be set as councilor. This address should be proposer address as well."
			}
		}
	}`)
}

func registerPermissionsCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgWhitelistPermissions{}, "tsukiHub/MsgWhitelistPermissions", nil)
	functionmeta.AddNewFunction((&MsgWhitelistPermissions{}).Type(), `{
		"description": "MsgWhitelistPermissions defines a message to whitelist permission of an address.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"address": {
				"type":        "string",
				"description": "Address to whitelist permission to."
			},
			"permission": {
				"type":        "uint32",
				"description": "Permission to be whitelisted."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgBlacklistPermissions{}, "tsukiHub/MsgBlacklistPermissions", nil)
	functionmeta.AddNewFunction((&MsgBlacklistPermissions{}).Type(), `{
		"description": "MsgBlacklistPermissions defines a message to blacklist permission of an address.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"address": {
				"type":        "string",
				"description": "Address to blacklist permission to."
			},
			"permission": {
				"type":        "uint32",
				"description": "Permission to be blacklisted."
			}
		}
	}`)
}

func registerRolesCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateRole{}, "tsukiHub/MsgCreateRole", nil)
	functionmeta.AddNewFunction((&MsgCreateRole{}).Type(), `{
		"description": "MsgCreateRole defines a message to create a role.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"role": {
				"type":        "uint32",
				"description": "Identifier of this role."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgAssignRole{}, "tsukiHub/MsgAssignRole", nil)
	functionmeta.AddNewFunction((&MsgAssignRole{}).Type(), `{
		"description": "MsgAssignRole defines a message to assign a role to an address.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"address": {
				"type":        "string",
				"description": "Address to set role to."
			},
			"role": {
				"type":        "uint32",
				"description": "role identifier."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgUnassignRole{}, "tsukiHub/MsgUnassignRole", nil)
	functionmeta.AddNewFunction((&MsgUnassignRole{}).Type(), `{
		"description": "MsgUnassignRole defines a message to unassign a role from an address.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"address": {
				"type":        "string",
				"description": "Address to remove role from."
			},
			"role": {
				"type":        "uint32",
				"description": "role identifier."
			}
		}
	}`)

	cdc.RegisterConcrete(&MsgWhitelistRolePermission{}, "tsukiHub/MsgWhitelistRolePermission", nil)
	functionmeta.AddNewFunction((&MsgWhitelistRolePermission{}).Type(), `{
		"description": "MsgWhitelistRolePermission defines a message to whitelist permission for a role.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"role": {
				"type":        "uint32",
				"description": "role identifier."
			},
			"permission": {
				"type":        "uint32",
				"description": "Permission to be whitelisted."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgBlacklistRolePermission{}, "tsukiHub/MsgBlacklistRolePermission", nil)
	functionmeta.AddNewFunction((&MsgBlacklistRolePermission{}).Type(), `{
		"description": "MsgBlacklistRolePermission defines a message to blacklist permission for a role.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"role": {
				"type":        "uint32",
				"description": "role identifier."
			},
			"permission": {
				"type":        "uint32",
				"description": "Permission to be blacklisted."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgRemoveWhitelistRolePermission{}, "tsukiHub/MsgRemoveWhitelistRolePermission", nil)
	functionmeta.AddNewFunction((&MsgRemoveWhitelistRolePermission{}).Type(), `{
		"description": "MsgRemoveWhitelistRolePermission defines a message to remove whitelisted permission for a role.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"role": {
				"type":        "uint32",
				"description": "role identifier."
			},
			"permission": {
				"type":        "uint32",
				"description": "Permission to be removed from whitelisted listing."
			}
		}
	}`)
	cdc.RegisterConcrete(&MsgRemoveBlacklistRolePermission{}, "tsukiHub/MsgRemoveBlacklistRolePermission", nil)
	functionmeta.AddNewFunction((&MsgRemoveBlacklistRolePermission{}).Type(), `{
		"description": "MsgRemoveBlacklistRolePermission defines a message to remove blacklisted permission for a role.",
		"parameters": {
			"proposer": {
				"type":        "string",
				"description": "proposer who propose this message."
			},
			"role": {
				"type":        "uint32",
				"description": "role identifier."
			},
			"permission": {
				"type":        "uint32",
				"description": "Permission to be removed from blacklisted listing."
			}
		}
	}`)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWhitelistPermissions{},
		&MsgBlacklistPermissions{},

		&MsgSetNetworkProperties{},
		&MsgSetExecutionFee{},

		&MsgClaimCouncilor{},

		&MsgAssignRole{},
		&MsgUnassignRole{},
		&MsgCreateRole{},

		&MsgWhitelistRolePermission{},
		&MsgBlacklistRolePermission{},
		&MsgRemoveWhitelistRolePermission{},
		&MsgRemoveBlacklistRolePermission{},

		&MsgSubmitProposal{},
		&MsgVoteProposal{},

		&MsgRegisterIdentityRecords{},
		&MsgDeleteIdentityRecords{},
		&MsgRequestIdentityRecordsVerify{},
		&MsgHandleIdentityRecordsVerifyRequest{},
		&MsgCancelIdentityRecordsVerifyRequest{},
	)

	registry.RegisterInterface(
		"tsuki.gov.Content",
		(*Content)(nil),
		&WhitelistAccountPermissionProposal{},
		&BlacklistAccountPermissionProposal{},
		&RemoveWhitelistedAccountPermissionProposal{},
		&RemoveBlacklistedAccountPermissionProposal{},
		&AssignRoleToAccountProposal{},
		&UnassignRoleFromAccountProposal{},
		&SetNetworkPropertyProposal{},
		&UpsertDataRegistryProposal{},
		&SetPoorNetworkMessagesProposal{},
		&CreateRoleProposal{},
		&RemoveRoleProposal{},
		&WhitelistRolePermissionProposal{},
		&BlacklistRolePermissionProposal{},
		&RemoveWhitelistedRolePermissionProposal{},
		&RemoveBlacklistedRolePermissionProposal{},
		&SetProposalDurationsProposal{},
		&ProposalResetWholeCouncilorRank{},
		&ProposalJailCouncilor{},
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
