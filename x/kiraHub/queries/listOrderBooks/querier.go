package listOrderBooks

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/cosmos-sdk/x/mint"
	"github.com/TsukiCore/tsuki/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func QueryGetOrderBooks(ctx sdk.Context, path []string, req abci.RequestQuery, keeper mint.Keeper) ([]byte, error) {

	var queryOutput []types.OrderBook

	if path[0] == "ID" {

		queryOutput = keeper.GetOrderBookByID(ctx, path[1])

	} else if path[0] == "Index" {

		queryOutput = keeper.GetOrderBookByIndex(ctx, path[1])

	} else if path[0] == "Quote" {

		queryOutput = keeper.GetOrderBooksByQuote(ctx, path[1])

	} else if path[0] == "Base" {

		queryOutput = keeper.GetOrderBooksByBase(ctx, path[1])

	} else if path[0] == "Trading_Pair" {

		queryOutput = keeper.GetOrderBooksByTP(ctx, path[1])

	} else if path[0] == "Curator" {

		queryOutput = keeper.GetOrderBooksByCurator(ctx, path[1])
	}

	res, marshalJSONIndentError := codec.MarshalJSONIndent(packageCodec, queryOutput)
	if marshalJSONIndentError != nil {
		panic(marshalJSONIndentError)
	}

	return res, nil


}