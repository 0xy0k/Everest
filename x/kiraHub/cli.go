package tsukiHub


import (
	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
	"github.com/spf13/cobra"

	"github.com/TsukiCore/cosmos-sdk/client"
	"github.com/TsukiCore/cosmos-sdk/client/flags"
	"github.com/TsukiCore/cosmos-sdk/codec"
)

func GetCLIRootTransactionCommand(codec *codec.Codec) *cobra.Command {
	rootTransactionCommand := &cobra.Command{
		Use:                        "testtoken",
		Short:                      "Asset root transaction command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	rootTransactionCommand.AddCommand(flags.PostCommands(
	)...)
	return rootTransactionCommand
}

func GetCLIRootQueryCommand(codec *codec.Codec) *cobra.Command {
	rootQueryCommand := &cobra.Command{
		Use:                        constants.QuerierRoute,
		Short:                      "Asset root query command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	rootQueryCommand.AddCommand(flags.GetCommands(
	)...)
	return rootQueryCommand
}