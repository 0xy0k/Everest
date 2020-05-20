package tsukiHub

import (
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/cosmos-sdk/types/errors"
	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
	"github.com/TsukiCore/tsuki/x/tsukiHub/queries/listOrderBooks"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryListOrderBooks = "listOrderBooks"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(context sdk.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		switch path[0] {

		case QueryListOrderBooks:
			return listOrderBooks.QueryGetOrderBooks(context, path[1:], requestQuery, keeper.getCreateOrderBookKeeper())

		default:
			return nil, errors.Wrapf(constants.UnknownQueryCode, "%v", path[0])
		}
	}
}
