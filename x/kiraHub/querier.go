package tsukiHub

import (
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
	"github.com/TsukiCore/cosmos-sdk/types/errors"
	abciTypes "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(context sdk.Context, path []string, requestQuery abciTypes.RequestQuery) ([]byte, error) {
		switch path[0] {


		default:
			return nil, errors.Wrapf(constants.UnknownQueryCode, "%v", path[0])
		}
	}
}
