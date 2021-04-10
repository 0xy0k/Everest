package types

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
)

// DefaultGenesis returns the default CustomGo genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Permissions: map[uint64]*Permissions{
			uint64(RoleSudo): NewPermissions([]PermValue{
				PermSetPermissions,
				PermClaimValidator,
				PermClaimCouncilor,
				PermUpsertTokenAlias,
				// PermChangeTxFee, // do not give this permission to sudo account - test does not pass
				PermUpsertTokenRate,
				PermUpsertRole,
				PermCreateSetPermissionsProposal,
				PermVoteSetPermissionProposal,
				PermCreateSetNetworkPropertyProposal,
				PermVoteSetNetworkPropertyProposal,
				PermCreateUpsertDataRegistryProposal,
				PermVoteUpsertDataRegistryProposal,
				PermCreateUpsertTokenAliasProposal,
				PermVoteUpsertTokenAliasProposal,
				PermCreateUpsertTokenRateProposal,
				PermVoteUpsertTokenRateProposal,
				PermCreateUnjailValidatorProposal,
				PermVoteUnjailValidatorProposal,
				PermCreateRoleProposal,
				PermVoteCreateRoleProposal,
				PermCreateTokensWhiteBlackChangeProposal,
				PermVoteTokensWhiteBlackChangeProposal,
				PermCreateSetPoorNetworkMessagesProposal,
				PermVoteSetPoorNetworkMessagesProposal,
			}, nil),
			uint64(RoleValidator): NewPermissions([]PermValue{PermClaimValidator}, nil),
		},
		StartingProposalId: 1,
		NetworkProperties: &NetworkProperties{
			MinTxFee:                    100,
			MaxTxFee:                    1000000,
			VoteQuorum:                  33,
			ProposalEndTime:             600, // 600 seconds / 10 mins
			ProposalEnactmentTime:       300, // 300 seconds / 5 mins
			MinProposalEndBlocks:        2,
			MinProposalEnactmentBlocks:  1,
			EnableForeignFeePayments:    true,
			MischanceRankDecreaseAmount: 10,
			InactiveRankDecreasePercent: 50,      // 50%
			PoorNetworkMaxBankSend:      1000000, // 1M ukex
			MinValidators:               1,
			JailMaxTime:                 10, // 10 mins
			EnableTokenWhitelist:        false,
			EnableTokenBlacklist:        true,
		},
		ExecutionFees: []*ExecutionFee{
			{
				Name:              "Claim Validator Seat",
				TransactionType:   "claim-validator-seat",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Claim Governance Seat",
				TransactionType:   "claim-governance-seat",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Claim Proposal Type X",
				TransactionType:   "claim-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Vote Proposal Type X",
				TransactionType:   "vote-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Submit Proposal Type X",
				TransactionType:   "submit-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Veto Proposal Type X",
				TransactionType:   "veto-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Upsert Token Alias Execution Fee",
				TransactionType:   tsukitypes.MsgTypeUpsertTokenAlias,
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Activate a validator",
				TransactionType:   tsukitypes.MsgTypeActivate,
				ExecutionFee:      100,
				FailureFee:        1000,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Pause a validator",
				TransactionType:   tsukitypes.MsgTypePause,
				ExecutionFee:      10,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				Name:              "Unpause a validator",
				TransactionType:   tsukitypes.MsgTypeUnpause,
				ExecutionFee:      10,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
		},
		PoorNetworkMessages: &AllowedMessages{
			Messages: []string{
				tsukitypes.MsgTypeProposalAssignPermission,
				tsukitypes.MsgTypeProposalSetNetworkProperty,
				tsukitypes.MsgTypeSetNetworkProperties,
				tsukitypes.MsgTypeVoteProposal,
				tsukitypes.MsgTypeClaimCouncilor,
				tsukitypes.MsgTypeWhitelistPermissions,
				tsukitypes.MsgTypeBlacklistPermissions,
				tsukitypes.MsgTypeCreateRole,
				tsukitypes.MsgTypeAssignRole,
				tsukitypes.MsgTypeRemoveRole,
				tsukitypes.MsgTypeWhitelistRolePermission,
				tsukitypes.MsgTypeBlacklistRolePermission,
				tsukitypes.MsgTypeRemoveWhitelistRolePermission,
				tsukitypes.MsgTypeRemoveBlacklistRolePermission,
				tsukitypes.MsgTypeClaimValidator,
				tsukitypes.MsgTypeActivate,
				tsukitypes.MsgTypePause,
				tsukitypes.MsgTypeUnpause,
			},
		},
	}
}
