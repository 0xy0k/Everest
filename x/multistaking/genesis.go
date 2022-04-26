package multistaking

import (
	"github.com/TsukiCore/tsuki/x/multistaking/keeper"
	"github.com/TsukiCore/tsuki/x/multistaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the evidence module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, gs *types.GenesisState) {

}

func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{}
}
