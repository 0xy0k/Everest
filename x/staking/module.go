package staking

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/cosmos-sdk/types/module"
	"github.com/TsukiCore/cosmos-sdk/x/staking"
	stakingkeeper "github.com/TsukiCore/cosmos-sdk/x/staking/keeper"
	"github.com/TsukiCore/cosmos-sdk/x/staking/types"
	"github.com/TsukiCore/tsuki/x/staking/keeper"
)

var (
	_ module.AppModule = AppModule{}
)

// AppModule extends the cosmos SDK staking.
type AppModule struct {
	staking.AppModule

	stakingKeeper       stakingkeeper.Keeper
	customStakingKeeper keeper.Keeper
}

// NewHandler returns an sdk.Handler for the staking module.
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.stakingKeeper, am.customStakingKeeper)
}

// NewAppModule returns a new Custom Staking module.
func NewAppModule(
	cdc codec.Marshaler,
	keeper stakingkeeper.Keeper,
	ak types.AccountKeeper,
	bk types.BankKeeper,
) AppModule {
	return AppModule{
		AppModule:     staking.NewAppModule(cdc, keeper, ak, bk),
		stakingKeeper: keeper,
	}
}
