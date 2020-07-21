package staking

import (
	"github.com/TsukiCore/cosmos-sdk/codec"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/cosmos-sdk/types/module"
	"github.com/TsukiCore/cosmos-sdk/x/auth"
	"github.com/TsukiCore/cosmos-sdk/x/staking"
	"github.com/TsukiCore/cosmos-sdk/x/staking/types"
)

var (
	_ module.AppModule = AppModule{}
)

// AppModule extends the cosmos SDK staking.
type AppModule struct {
	staking.AppModule

	stakingKeeper staking.Keeper
}

// NewHandler returns an sdk.Handler for the staking module.
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.stakingKeeper)
}

// NewAppModule returns a new Custom Staking module.
func NewAppModule(
	cdc codec.Marshaler,
	keeper staking.Keeper,
	ak auth.AccountKeeper,
	bk types.BankKeeper,
) AppModule {
	return AppModule{
		AppModule:     staking.NewAppModule(cdc, keeper, ak, bk),
		stakingKeeper: keeper,
	}
}
