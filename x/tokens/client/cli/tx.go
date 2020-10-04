package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/TsukiCore/tsuki/x/tokens/types"
)

// flags for tokens module txs
const (
	FlagExpiration       = "expiration"
	FlagEnactment        = "enactment"
	FlagAllowedVoteTypes = "allowed_vote_types"
	FlagSymbol           = "symbol"
	FlagName             = "name"
	FlagIcon             = "icon"
	FlagDecimals         = "decimals"
	FlagDenoms           = "denoms"
)

// NewTxCmd returns a root CLI command handler for all x/bank transaction commands.
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Tokens sub commands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(GetTxUpsertTokenAliasCmd())

	return txCmd
}

// GetTxUpsertTokenAliasCmd implement cli command for MsgUpsertTokenAlias
func GetTxUpsertTokenAliasCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upsert-alias",
		Short: "Upsert token alias",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			expiration, err := cmd.Flags().GetUint32(FlagExpiration)
			if err != nil {
				return fmt.Errorf("invalid expiration")
			}

			enactment, err := cmd.Flags().GetUint32(FlagEnactment)
			if err != nil {
				return fmt.Errorf("invalid enactment")
			}

			allowedVoteTypesString, err := cmd.Flags().GetString(FlagAllowedVoteTypes)
			if err != nil {
				return fmt.Errorf("invalid vote type")
			}

			allowedVoteTypes := []types.VoteType{}
			for _, v := range strings.Split(allowedVoteTypesString, ",") {
				voteType, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("invalid vote type")
				}
				allowedVoteTypes = append(allowedVoteTypes, types.VoteType(voteType))
			}

			symbol, err := cmd.Flags().GetString(FlagSymbol)
			if err != nil {
				return fmt.Errorf("invalid symbol")
			}

			name, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return fmt.Errorf("invalid name")
			}

			icon, err := cmd.Flags().GetString(FlagIcon)
			if err != nil {
				return fmt.Errorf("invalid icon")
			}

			decimals, err := cmd.Flags().GetUint32(FlagDecimals)
			if err != nil {
				return fmt.Errorf("invalid decimals")
			}

			denomsString, err := cmd.Flags().GetString(FlagDenoms)
			if err != nil {
				return fmt.Errorf("invalid denoms")
			}

			denoms := strings.Split(denomsString, ",")
			// TODO should do validation for denom regex

			status := types.ProposalStatus_undefined

			msg := types.NewMsgUpsertTokenAlias(
				clientCtx.FromAddress,
				expiration,
				enactment,
				allowedVoteTypes,
				symbol,
				name,
				icon,
				decimals,
				denoms,
				status,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint32(FlagExpiration, 0, "expiration time - seconds since submission")
	cmd.Flags().Uint32(FlagEnactment, 0, "enactment time - seconds since expiration")
	cmd.Flags().String(FlagAllowedVoteTypes, "0,1,2,3", "Allowed vote types: yes(0), no(1), veto(2), abstain(3)")
	cmd.Flags().String(FlagSymbol, "KEX", "Ticker (eg. ATOM, KEX, BTC)")
	cmd.Flags().String(FlagName, "Tsuki", "Token Name (e.g. Cosmos, Tsuki, Bitcoin)")
	cmd.Flags().String(FlagIcon, "", "Graphical Symbol (url link to graphics)")
	cmd.Flags().Uint32(FlagDecimals, 6, "Integer number of max decimals")
	cmd.Flags().String(FlagDenoms, "ukex,mkex", "An array of token denoms to be aliased")

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	return cmd
}
