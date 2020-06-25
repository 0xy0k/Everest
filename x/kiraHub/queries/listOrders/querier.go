package listOrders

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrder"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrderBook"
	abci "github.com/tendermint/tendermint/abci/types"
	"strconv"
)

func QueryGetOrders(ctx sdk.Context, path []string, req abci.RequestQuery, keeper createOrder.Keeper) ([]byte, error) {

	var queryOutput []types.LimitOrder

	var int1, _ = strconv.Atoi(path[1])
	var int2, _ = strconv.Atoi(path[2])

	queryOutput = keeper.GetOrders(ctx, path[0], uint32(int1), uint32(int2))

	res, marshalJSONIndentError := codec.MarshalJSONIndent(packageCodec, queryOutput)
	if marshalJSONIndentError != nil {
		panic(marshalJSONIndentError)
	}

	return res, nil
}