package staking

import (
	"encoding/json"

	"github.com/TsukiCore/tsuki/x/staking/client/cli"

	"github.com/TsukiCore/cosmos-sdk/client"
	"github.com/TsukiCore/cosmos-sdk/codec"
	types2 "github.com/TsukiCore/cosmos-sdk/codec/types"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/cosmos-sdk/types/module"
	"github.com/TsukiCore/cosmos-sdk/x/staking"
	stakingkeeper "github.com/TsukiCore/cosmos-sdk/x/staking/keeper"
	"github.com/TsukiCore/cosmos-sdk/x/staking/types"
	"github.com/TsukiCore/tsuki/x/staking/keeper"
	cumstomtypes "github.com/TsukiCore/tsuki/x/staking/types"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

type AppModuleBasic struct {
}

func (b AppModuleBasic) Name() string {
	return cumstomtypes.ModuleName
}

func (b AppModuleBasic) RegisterInterfaces(registry types2.InterfaceRegistry) {
	cumstomtypes.RegisterInterfaces(registry)
}

func (b AppModuleBasic) DefaultGenesis(marshaler codec.JSONMarshaler) json.RawMessage {
	return nil
}

func (b AppModuleBasic) ValidateGenesis(marshaler codec.JSONMarshaler, config client.TxEncodingConfig, message json.RawMessage) error {
	return nil
}

func (b AppModuleBasic) RegisterRESTRoutes(context client.Context, router *mux.Router) {
}

func (b AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxClaimValidatorCmd()
}

func (b AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil
}

// RegisterCodec registers the staking module's types for the given codec.
func (AppModuleBasic) RegisterCodec(cdc *codec.Codec) {
	cumstomtypes.RegisterCodec(cdc)
}

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
