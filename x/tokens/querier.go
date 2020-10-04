package tokens

import (
	"github.com/TsukiCore/tsuki/x/tokens/keeper"
	"github.com/TsukiCore/tsuki/x/tokens/types"
)

type Querier struct {
	keeper keeper.Keeper
}

func NewQuerier(keeper keeper.Keeper) types.QueryServer {
	return &Querier{keeper: keeper}
}
