package cli

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/TsukiCore/tsuki/x/genutil"
	v01228govtypes "github.com/TsukiCore/tsuki/x/gov/legacy/v01228"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	upgradetypes "github.com/TsukiCore/tsuki/x/upgrade/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
)

// GetNewGenesisFromExportedCmd returns new genesis from exported genesis
func GetNewGenesisFromExportedCmd(mbm module.BasicManager, txEncCfg client.TxEncodingConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "new-genesis-from-exported [path-to-exported.json] [path-to-new.json]",
		Short: "Get new genesis from exported app state json",
		Args:  cobra.ExactArgs(2),
		Long: fmt.Sprintf(`Get new genesis from exported app state json.
- Change chain-id to new_chain_id as indicated by the upgrade plan
- Replace current upgrade plan in the app_state.upgrade with next plan and set next plan to null

Example:
$ %s new-genesis-from-exported exported-genesis.json new-genesis.json
`, version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			cdc := clientCtx.JSONCodec

			genDoc, err := tmtypes.GenesisDocFromFile(args[0])
			if err != nil {
				return errors.Wrapf(err, "failed to read genesis file %s", args[0])
			}

			var genesisState map[string]json.RawMessage
			if err = json.Unmarshal(genDoc.AppState, &genesisState); err != nil {
				return errors.Wrap(err, "failed to unmarshal genesis state")
			}

			if err = mbm.ValidateGenesis(cdc, txEncCfg, genesisState); err != nil {
				return errors.Wrap(err, "failed to validate genesis state")
			}

			upgradeGenesis := upgradetypes.GenesisState{}
			cdc.MustUnmarshalJSON(genesisState[upgradetypes.ModuleName], &upgradeGenesis)
			oldVersion := upgradeGenesis.Version
			if upgradeGenesis.Version == "" {
				upgradeGenesis.Version = "v0.1.22.11"
				fmt.Println("upgraded the upgrade module genesis to v0.1.22.11")
			}

			if upgradeGenesis.NextPlan == nil {
				return fmt.Errorf("next plan is not available")
			}

			if genDoc.ChainID != upgradeGenesis.NextPlan.OldChainId {
				return fmt.Errorf("next plan has different oldchain id, current chain_id=%s, next_plan.old_chain_id=%s", genDoc.ChainID, upgradeGenesis.NextPlan.OldChainId)
			}

			genDoc.ChainID = upgradeGenesis.NextPlan.NewChainId
			upgradeGenesis.CurrentPlan = upgradeGenesis.NextPlan
			upgradeGenesis.NextPlan = nil

			genesisState[upgradetypes.ModuleName] = cdc.MustMarshalJSON(&upgradeGenesis)

			govGenesisV01228 := v01228govtypes.GenesisStateV01228{}
			err = cdc.UnmarshalJSON(genesisState[govtypes.ModuleName], &govGenesisV01228)

			// we are referencing oldPlan.name to determine upgrade genesis or not
			if err == nil && oldVersion == "" { // it means v0.1.22.8 gov genesis
				govGenesis := govtypes.GenesisState{
					StartingProposalId: govGenesisV01228.StartingProposalId,
					NextRoleId:         govtypes.DefaultGenesis().NextRoleId,
					Roles:              govtypes.DefaultGenesis().Roles,
					RolePermissions:    govGenesisV01228.Permissions,
					NetworkActors:      govGenesisV01228.NetworkActors,
					NetworkProperties: &govtypes.NetworkProperties{
						MinTxFee:                    govGenesisV01228.NetworkProperties.MinTxFee,
						MaxTxFee:                    govGenesisV01228.NetworkProperties.MaxTxFee,
						VoteQuorum:                  govGenesisV01228.NetworkProperties.VoteQuorum,
						MinimumProposalEndTime:      govGenesisV01228.NetworkProperties.ProposalEndTime,
						ProposalEnactmentTime:       govGenesisV01228.NetworkProperties.ProposalEnactmentTime,
						MinProposalEndBlocks:        govGenesisV01228.NetworkProperties.MinProposalEndBlocks,
						MinProposalEnactmentBlocks:  govGenesisV01228.NetworkProperties.MinProposalEnactmentBlocks,
						EnableForeignFeePayments:    govGenesisV01228.NetworkProperties.EnableForeignFeePayments,
						MischanceRankDecreaseAmount: govGenesisV01228.NetworkProperties.MischanceRankDecreaseAmount,
						MaxMischance:                govGenesisV01228.NetworkProperties.MaxMischance,
						MischanceConfidence:         govGenesisV01228.NetworkProperties.MischanceConfidence,
						InactiveRankDecreasePercent: govGenesisV01228.NetworkProperties.InactiveRankDecreasePercent,
						MinValidators:               govGenesisV01228.NetworkProperties.MinValidators,
						PoorNetworkMaxBankSend:      govGenesisV01228.NetworkProperties.PoorNetworkMaxBankSend,
						UnjailMaxTime:               govGenesisV01228.NetworkProperties.JailMaxTime,
						EnableTokenWhitelist:        govGenesisV01228.NetworkProperties.EnableTokenWhitelist,
						EnableTokenBlacklist:        govGenesisV01228.NetworkProperties.EnableTokenBlacklist,
						MinIdentityApprovalTip:      govGenesisV01228.NetworkProperties.MinIdentityApprovalTip,
						UniqueIdentityKeys:          govGenesisV01228.NetworkProperties.UniqueIdentityKeys,
						UbiHardcap:                  6000_000,
						ValidatorsFeeShare:          50,
						InflationRate:               18,       // 18%
						InflationPeriod:             31557600, // 1 year
						UnstakingPeriod:             2629800,  // 1 month
						MaxDelegators:               100,
						MinDelegationPushout:        10,
						MinCustodyReward:            200,
						MaxCustodyTxSize:            8192,
						MaxCustodyBufferSize:        10,
					},
					ExecutionFees:               govGenesisV01228.ExecutionFees,
					PoorNetworkMessages:         govGenesisV01228.PoorNetworkMessages,
					Proposals:                   []govtypes.Proposal{}, // govGenesisV01228.Proposals,
					Votes:                       []govtypes.Vote{},     // govGenesisV01228.Votes,
					DataRegistry:                govGenesisV01228.DataRegistry,
					IdentityRecords:             govGenesisV01228.IdentityRecords,
					LastIdentityRecordId:        govGenesisV01228.LastIdentityRecordId,
					IdRecordsVerifyRequests:     govGenesisV01228.IdRecordsVerifyRequests,
					LastIdRecordVerifyRequestId: govGenesisV01228.LastIdRecordVerifyRequestId,
					ProposalDurations:           make(map[string]uint64),
				}

				genesisState[govtypes.ModuleName] = cdc.MustMarshalJSON(&govGenesis)
			} else {
				fmt.Println("GovGenesis01228 unmarshal test: ", err)
				fmt.Println("Skipping governance module upgrade since it is not v0.1.22.8 genesis")
			}

			// upgrade gov genesis for more role permissions
			govGenesis := govtypes.GenesisState{}
			err = cdc.UnmarshalJSON(genesisState[govtypes.ModuleName], &govGenesis)
			if err == nil {
				govGenesis.RolePermissions[govtypes.RoleSudo] = govtypes.DefaultGenesis().RolePermissions[govtypes.RoleSudo]
				genesisState[govtypes.ModuleName] = cdc.MustMarshalJSON(&govGenesis)
			}

			appState, err := json.MarshalIndent(genesisState, "", " ")
			if err != nil {
				return errors.Wrap(err, "Failed to marshall default genesis state")
			}

			genDoc.AppState = appState
			if err = genutil.ExportGenesisFile(genDoc, args[1]); err != nil {
				return errors.Wrap(err, "Failed to export genesis file")
			}
			return nil
		},
	}

	return cmd
}
