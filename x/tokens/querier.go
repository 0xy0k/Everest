package tokens

import (
	"context"

	"github.com/TsukiCore/tsuki/x/tokens/keeper"
	"github.com/TsukiCore/tsuki/x/tokens/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Querier struct {
	keeper keeper.Keeper
}

func NewQuerier(keeper keeper.Keeper) types.QueryServer {
	return &Querier{keeper: keeper}
}

func (q Querier) GetTokenAlias(ctx context.Context, request *types.TokenAliasRequest) (*types.TokenAliasResponse, error) {
	alias := q.keeper.GetTokenAlias(sdk.UnwrapSDKContext(ctx), request.Symbol)

	return &types.TokenAliasResponse{Data: alias}, nil
}
