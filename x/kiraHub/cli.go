package tsukiHub


import (
	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
	"github.com/TsukiCore/tsuki/x/tsukiHub/queries/listOrderBooks"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrderBook"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrder"
	"github.com/spf13/cobra"

	"github.com/TsukiCore/cosmos-sdk/client"
	"github.com/TsukiCore/cosmos-sdk/client/flags"
	"github.com/TsukiCore/cosmos-sdk/codec"
)

func GetCLIRootTransactionCommand(codec *codec.Codec) *cobra.Command {
	rootTransactionCommand := &cobra.Command{
		Use:                        constants.TransactionRoute,
		Short:                      "Asset root transaction command.",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	rootTransactionCommand.AddCommand(flags.PostCommands(
		createOrderBook.TransactionCommand(codec),
		createOrder.TransactionCommand(codec),
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
		listOrderBooks.GetOrderBooksCmd(codec),
		listOrderBooks.GetOrderBooksByTPCmd(codec),
	)...)
	return rootQueryCommand
}