package types

import (
	"encoding/json"

	appparams "github.com/TsukiCore/tsuki/app/params"
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DefaultGenesis returns the default CustomGo genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DefaultDenom: appparams.DefaultDenom,
		Bech32Prefix: appparams.AccountAddressPrefix,
		NextRoleId:   3,
		Roles: []Role{
			{
				Id:          uint32(RoleSudo),
				Sid:         "sudo",
				Description: "Sudo role",
			},
			{
				Id:          uint32(RoleValidator),
				Sid:         "validator",
				Description: "Validator role",
			},
		},
		RolePermissions: map[uint64]*Permissions{
			RoleSudo: NewPermissions([]PermValue{
				PermSetPermissions,
				PermClaimValidator,
				PermClaimCouncilor,
				PermUpsertTokenAlias,
				// PermChangeTxFee, // do not give this permission to sudo account - test does not pass
				PermUpsertTokenRate,
				PermUpsertRole,
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
				PermCreateSetProposalDurationProposal,
				PermVoteSetProposalDurationProposal,
				PermCreateTokensWhiteBlackChangeProposal,
				PermVoteTokensWhiteBlackChangeProposal,
				PermCreateSetPoorNetworkMessagesProposal,
				PermVoteSetPoorNetworkMessagesProposal,
				PermWhitelistAccountPermissionProposal,
				PermVoteWhitelistAccountPermissionProposal,
				PermCreateResetWholeValidatorRankProposal,
				PermVoteResetWholeValidatorRankProposal,
				PermCreateSoftwareUpgradeProposal,
				PermVoteSoftwareUpgradeProposal,
				PermSetClaimValidatorPermission,
				PermBlacklistAccountPermissionProposal,
				PermVoteBlacklistAccountPermissionProposal,
				PermRemoveWhitelistedAccountPermissionProposal,
				PermVoteRemoveWhitelistedAccountPermissionProposal,
				PermRemoveBlacklistedAccountPermissionProposal,
				PermVoteRemoveBlacklistedAccountPermissionProposal,
				PermWhitelistRolePermissionProposal,
				PermVoteWhitelistRolePermissionProposal,
				PermBlacklistRolePermissionProposal,
				PermVoteBlacklistRolePermissionProposal,
				PermRemoveWhitelistedRolePermissionProposal,
				PermVoteRemoveWhitelistedRolePermissionProposal,
				PermRemoveBlacklistedRolePermissionProposal,
				PermVoteRemoveBlacklistedRolePermissionProposal,
				PermAssignRoleToAccountProposal,
				PermVoteAssignRoleToAccountProposal,
				PermUnassignRoleFromAccountProposal,
				PermVoteUnassignRoleFromAccountProposal,
				PermRemoveRoleProposal,
				PermVoteRemoveRoleProposal,
				PermCreateUpsertUBIProposal,
				PermVoteUpsertUBIProposal,
				PermCreateRemoveUBIProposal,
				PermVoteRemoveUBIProposal,
				PermCreateSlashValidatorProposal,
				PermVoteSlashValidatorProposal,
				PermCreateBasketProposal,
				PermVoteBasketProposal,
				PermHandleBasketEmergency,
				PermCreateResetWholeCouncilorRankProposal,
				PermVoteResetWholeCouncilorRankProposal,
				PermCreateJailCouncilorProposal,
				PermVoteJailCouncilorProposal,
				PermCreatePollProposal,
			}, nil),
			uint64(RoleValidator): NewPermissions([]PermValue{PermClaimValidator}, nil),
		},
		StartingProposalId: 1,
		NetworkProperties: &NetworkProperties{
			MinTxFee:                        100,
			MaxTxFee:                        1000000,
			VoteQuorum:                      33,
			MinimumProposalEndTime:          300, // 300 seconds / 5 mins
			ProposalEnactmentTime:           300, // 300 seconds / 5 mins
			MinProposalEndBlocks:            2,
			MinProposalEnactmentBlocks:      1,
			EnableForeignFeePayments:        true,
			MischanceRankDecreaseAmount:     10,
			MischanceConfidence:             10,
			MaxMischance:                    110,
			InactiveRankDecreasePercent:     sdk.NewDecWithPrec(50, 2), // 50%
			PoorNetworkMaxBankSend:          1000000,                   // 1M ukex
			MinValidators:                   1,
			UnjailMaxTime:                   600, // 600  seconds / 10 mins
			EnableTokenWhitelist:            false,
			EnableTokenBlacklist:            true,
			MinIdentityApprovalTip:          200,
			UniqueIdentityKeys:              "moniker,username",
			UbiHardcap:                      6000_000,
			ValidatorsFeeShare:              sdk.NewDecWithPrec(50, 2), // 50%
			InflationRate:                   sdk.NewDecWithPrec(18, 2), // 18%
			InflationPeriod:                 31557600,                  // 1 year
			UnstakingPeriod:                 2629800,                   // 1 month
			MaxDelegators:                   100,
			MinDelegationPushout:            10,
			SlashingPeriod:                  3600,
			MaxJailedPercentage:             sdk.NewDecWithPrec(25, 2),
			MaxSlashingPercentage:           sdk.NewDecWithPrec(1, 2),
			MinCustodyReward:                200,
			MaxCustodyTxSize:                8192,
			MaxCustodyBufferSize:            10,
			AbstentionRankDecreaseAmount:    1,
			MaxAbstention:                   2,
			MinCollectiveBond:               100_000, // in KEX
			MinCollectiveBondingTime:        86400,   // in seconds
			MaxCollectiveOutputs:            10,
			MinCollectiveClaimPeriod:        14400,                     // 4hrs
			ValidatorRecoveryBond:           300000,                    // 300k KEX
			MaxAnnualInflation:              sdk.NewDecWithPrec(35, 2), // 35%// 300k KEX
			MaxProposalTitleSize:            128,
			MaxProposalDescriptionSize:      1024,
			MaxProposalPollOptionSize:       64,
			MaxProposalPollOptionCount:      128,
			MaxProposalReferenceSize:        512,
			MaxProposalChecksumSize:         128,
			MinDappBond:                     1000000,
			MaxDappBond:                     10000000,
			DappBondDuration:                604800,
			DappVerifierBond:                sdk.NewDecWithPrec(1, 3), //0.1%
			DappAutoDenounceTime:            60,                       // 60s
			DappMischanceRankDecreaseAmount: 1,
			DappMaxMischance:                10,
			DappInactiveRankDecreasePercent: 10,
			DappPoolSlippageDefault:         sdk.NewDecWithPrec(1, 1), // 10%
			MintingFtFee:                    100_000_000_000_000,
			MintingNftFee:                   100_000_000_000_000,
			VetoThreshold:                   sdk.NewDecWithPrec(3340, 2), // 33.40%
			AutocompoundIntervalNumBlocks:   17280,
		},
		ExecutionFees: []ExecutionFee{
			{
				TransactionType:   tsukitypes.MsgTypeClaimValidator,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypeClaimCouncilor,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "claim-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "vote-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "submit-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "veto-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypeUpsertTokenAlias,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypeActivate,
				ExecutionFee:      100,
				FailureFee:        1000,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypePause,
				ExecutionFee:      100,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   tsukitypes.MsgTypeUnpause,
				ExecutionFee:      100,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
		},
		PoorNetworkMessages: &AllowedMessages{
			Messages: []string{
				tsukitypes.MsgTypeSubmitProposal,
				tsukitypes.MsgTypeSetNetworkProperties,
				tsukitypes.MsgTypeVoteProposal,
				tsukitypes.MsgTypeClaimCouncilor,
				tsukitypes.MsgTypeWhitelistPermissions,
				tsukitypes.MsgTypeBlacklistPermissions,
				tsukitypes.MsgTypeCreateRole,
				tsukitypes.MsgTypeAssignRole,
				tsukitypes.MsgTypeUnassignRole,
				tsukitypes.MsgTypeWhitelistRolePermission,
				tsukitypes.MsgTypeBlacklistRolePermission,
				tsukitypes.MsgTypeRemoveWhitelistRolePermission,
				tsukitypes.MsgTypeRemoveBlacklistRolePermission,
				tsukitypes.MsgTypeClaimValidator,
				tsukitypes.MsgTypeActivate,
				tsukitypes.MsgTypePause,
				tsukitypes.MsgTypeUnpause,
				tsukitypes.MsgTypeRegisterIdentityRecords,
				tsukitypes.MsgTypeDeleteIdentityRecord,
				tsukitypes.MsgTypeRequestIdentityRecordsVerify,
				tsukitypes.MsgTypeHandleIdentityRecordsVerifyRequest,
				tsukitypes.MsgTypeCancelIdentityRecordsVerifyRequest,
			},
		},
		LastIdentityRecordId:        0,
		LastIdRecordVerifyRequestId: 0,
	}
}

// GetGenesisStateFromAppState returns x/auth GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.Codec, appState map[string]json.RawMessage) GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return genesisState
}

func GetBech32PrefixAndDefaultDenomFromAppState(appState map[string]json.RawMessage) (string, string) {
	var genesisState map[string]interface{}
	err := json.Unmarshal(appState[ModuleName], &genesisState)
	if err != nil {
		panic(err)
	}
	return genesisState["bech32_prefix"].(string), genesisState["default_denom"].(string)
}
