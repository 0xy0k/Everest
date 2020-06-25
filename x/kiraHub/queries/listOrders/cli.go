package listOrders

import (
	"fmt"
	"github.com/TsukiCore/cosmos-sdk/client/context"
	"github.com/TsukiCore/cosmos-sdk/codec"
	"github.com/TsukiCore/tsuki/types"

	"github.com/spf13/cobra"
)

func GetOrdersCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "listorders [id] [max_orders] [min_amount]",
		Short: "List order(s) by ID",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			//var owner = cliCtx.GetFromAddress()

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/tsukiHub/listOrders/%s/%s/%s", args[0], args[1], args[2]), nil)
			if err != nil {
				fmt.Printf("could not query. Searching By - %s with max_orders - %s and min_amount - %s \n", args[0], args[1], args[2])
				return nil
			}

			var out []types.LimitOrder
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
