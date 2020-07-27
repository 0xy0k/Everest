package app

import (
	"encoding/json"
	"io"
	"os"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	bam "github.com/TsukiCore/cosmos-sdk/baseapp"
	"github.com/TsukiCore/cosmos-sdk/codec"
	cdctypes "github.com/TsukiCore/cosmos-sdk/codec/types"
	"github.com/TsukiCore/cosmos-sdk/simapp"
	"github.com/TsukiCore/cosmos-sdk/std"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	"github.com/TsukiCore/cosmos-sdk/types/module"
	"github.com/TsukiCore/cosmos-sdk/version"
	"github.com/TsukiCore/cosmos-sdk/x/auth"
	"github.com/TsukiCore/cosmos-sdk/x/auth/ante"
	authkeeper "github.com/TsukiCore/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/TsukiCore/cosmos-sdk/x/auth/types"
	"github.com/TsukiCore/cosmos-sdk/x/bank"
	bankkeeper "github.com/TsukiCore/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/TsukiCore/cosmos-sdk/x/bank/types"
	"github.com/TsukiCore/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/TsukiCore/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/TsukiCore/cosmos-sdk/x/capability/types"
	"github.com/TsukiCore/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/TsukiCore/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/TsukiCore/cosmos-sdk/x/crisis/types"
	distr "github.com/TsukiCore/cosmos-sdk/x/distribution"
	distrclient "github.com/TsukiCore/cosmos-sdk/x/distribution/client"
	distrkeeper "github.com/TsukiCore/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/TsukiCore/cosmos-sdk/x/distribution/types"
	"github.com/TsukiCore/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/TsukiCore/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/TsukiCore/cosmos-sdk/x/evidence/types"
	"github.com/TsukiCore/cosmos-sdk/x/genutil"
	"github.com/TsukiCore/cosmos-sdk/x/gov"
	govkeeper "github.com/TsukiCore/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/TsukiCore/cosmos-sdk/x/gov/types"
	"github.com/TsukiCore/cosmos-sdk/x/ibc"
	ibctransferkeeper "github.com/TsukiCore/cosmos-sdk/x/ibc-transfer/keeper"
	ibctransfertypes "github.com/TsukiCore/cosmos-sdk/x/ibc-transfer/types"
	ibcclient "github.com/TsukiCore/cosmos-sdk/x/ibc/02-client"
	ibchost "github.com/TsukiCore/cosmos-sdk/x/ibc/24-host"
	ibckeeper "github.com/TsukiCore/cosmos-sdk/x/ibc/keeper"
	minttypes "github.com/TsukiCore/cosmos-sdk/x/mint/types"
	"github.com/TsukiCore/cosmos-sdk/x/params"
	paramsclient "github.com/TsukiCore/cosmos-sdk/x/params/client"
	paramskeeper "github.com/TsukiCore/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/TsukiCore/cosmos-sdk/x/params/types"
	paramproposal "github.com/TsukiCore/cosmos-sdk/x/params/types/proposal"
	"github.com/TsukiCore/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/TsukiCore/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/TsukiCore/cosmos-sdk/x/slashing/types"
	"github.com/TsukiCore/cosmos-sdk/x/staking"
	stakingkeeper "github.com/TsukiCore/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/TsukiCore/cosmos-sdk/x/staking/types"
	"github.com/TsukiCore/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/TsukiCore/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/TsukiCore/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/TsukiCore/cosmos-sdk/x/upgrade/types"
	port "github.com/cosmos/cosmos-sdk/x/ibc/05-port"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"

	"github.com/TsukiCore/tsuki/x/tsukiHub"
	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
)

const appName = "Tsuki"

var (

	// DefaultCLIHome default home directories for the application CLI
	DefaultCLIHome = os.ExpandEnv("$HOME/.tsukicli")

	// DefaultNodeHome sets the folder where the applcation data and configuration will be stored
	DefaultNodeHome = os.ExpandEnv("$HOME/.tsukid")

	// ModuleBasics The module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsclient.ProposalHandler, distrclient.ProposalHandler, upgradeclient.ProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},

		//tsukiHub.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distrtypes.ModuleName:          nil,
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
	}

	allowedReceivingModAcc = map[string]bool{
		distrtypes.ModuleName: true,
	}
)

// MakeCodec creates the application codec. The codec is sealed before it is
// returned.
func MakeCodec() (*std.Codec, *codec.Codec) {
	cdc := std.MakeCodec(ModuleBasics)
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	appCodec := std.NewAppCodec(cdc, interfaceRegistry)

	sdk.RegisterInterfaces(interfaceRegistry)
	ModuleBasics.RegisterInterfaceModules(interfaceRegistry)

	return appCodec, cdc
}

// NewApp extended ABCI application
type TsukiApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tKeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	accountKeeper    authkeeper.AccountKeeper
	bankKeeper       bankkeeper.Keeper
	tsukiHubKeeper    tsukiHub.Keeper
	capabilityKeeper *capabilitykeeper.Keeper
	stakingKeeper    stakingkeeper.Keeper
	slashingKeeper   slashingkeeper.Keeper
	distrKeeper      distrkeeper.Keeper
	govKeeper        govkeeper.Keeper
	crisisKeeper     crisiskeeper.Keeper
	upgradeKeeper    upgradekeeper.Keeper
	paramsKeeper     paramskeeper.Keeper
	ibcKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	evidenceKeeper   evidencekeeper.Keeper
	transferKeeper   ibctransferkeeper.Keeper

	// make scoped keepers public for test purposes
	scopedIBCKeeper      capabilitykeeper.ScopedKeeper
	scopedTransferKeeper capabilitykeeper.ScopedKeeper

	// Module Manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager
}

// verify app interface at compile time
var _ simapp.App = (*TsukiApp)(nil)

// NewhubApp is a constructor function for hubApp
func NewInitApp(
	logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, skipUpgradeHeights map[int64]bool, home string,
	baseAppOptions ...func(*bam.BaseApp),

) *TsukiApp {
	// TODO: Remove cdc in favor of appCodec once all modules are migrated.
	encodingConfig := MakeEncodingConfig()
	appCodec := encodingConfig.Marshaler
	cdc := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	// BaseApp handles interactions with Tendermint through the ABCI protocol
	bApp := bam.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.GRPCQueryRouter().SetAnyUnpacker(interfaceRegistry)

	// TODO: Add the keys that module requires
	keys := sdk.NewKVStoreKeys(authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey, constants.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// Here you initialize your application with the store keys it requires
	app := &TsukiApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		tKeys:          tKeys,
		memKeys:        memKeys,
	}

	app.paramsKeeper = initParamsKeeper(appCodec, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey])

	bApp.SetParamStore(app.paramsKeeper.Subspace(bam.Paramspace).WithKeyTable(std.ConsensusParamsKeyTable()))

	app.capabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])
	scopedIBCKeeper := app.capabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedTransferKeeper := app.capabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)

	// The AccountKeeper handles address -> account lookups
	app.accountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], app.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)
	app.bankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], app.accountKeeper, app.GetSubspace(banktypes.ModuleName), app.BlacklistedAccAddrs(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], app.accountKeeper, app.bankKeeper, app.GetSubspace(stakingtypes.ModuleName),
	)
	app.distrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.accountKeeper, app.bankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName, app.ModuleAccountAddrs(),
	)
	app.slashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, app.GetSubspace(slashingtypes.ModuleName),
	)
	app.crisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName), invCheckPeriod, app.bankKeeper, authtypes.FeeCollectorName,
	)
	app.upgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, home)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.distrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.upgradeKeeper))
	app.govKeeper = govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], app.GetSubspace(govtypes.ModuleName), app.accountKeeper, app.bankKeeper,
		&stakingKeeper, govRouter,
	)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.stakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(app.distrKeeper.Hooks(), app.slashingKeeper.Hooks()),
	)

	app.ibcKeeper = ibc.NewKeeper(
		app.cdc, appCodec, keys[ibc.StoreKey], app.stakingKeeper, scopedIBCKeeper,
	)

	// Create Transfer Keepers
	app.transferKeeper = transfer.NewKeeper(
		appCodec, keys[transfer.StoreKey],
		app.ibcKeeper.ChannelKeeper, &app.ibcKeeper.PortKeeper,
		app.accountKeeper, app.bankKeeper,
		scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.transferKeeper)

	app.tsukiHubKeeper = tsukiHub.NewKeeper(app.cdc, keys[constants.StoreKey])

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := port.NewRouter()
	ibcRouter.AddRoute(transfer.ModuleName, transferModule)
	app.ibcKeeper.SetRouter(ibcRouter)

	// create evidence keeper with router
	evidenceKeeper := evidence.NewKeeper(
		appCodec, keys[evidence.StoreKey], &app.stakingKeeper, app.slashingKeeper,
	)
	evidenceRouter := evidence.NewRouter().
		AddRoute(ibcclient.RouterKey, ibcclient.HandlerClientMisbehaviour(app.ibcKeeper.ClientKeeper))

	evidenceKeeper.SetRouter(evidenceRouter)
	app.evidenceKeeper = *evidenceKeeper

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(appCodec, app.accountKeeper),
		tsukiHub.NewAppModule(app.tsukiHubKeeper),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper),
		capability.NewAppModule(appCodec, *app.capabilityKeeper),
		crisis.NewAppModule(&app.crisisKeeper),
		gov.NewAppModule(appCodec, app.govKeeper, app.accountKeeper, app.bankKeeper),
		slashing.NewAppModule(appCodec, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		distr.NewAppModule(appCodec, app.distrKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		staking.NewAppModule(appCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		ibc.NewAppModule(app.ibcKeeper),
		params.NewAppModule(app.paramsKeeper),
		transferModule,
	)
	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.

	app.mm.SetOrderBeginBlockers(
		upgrade.ModuleName, distr.ModuleName, slashing.ModuleName,
		evidence.ModuleName, staking.ModuleName, ibc.ModuleName,
	)
	app.mm.SetOrderEndBlockers(crisis.ModuleName, gov.ModuleName, staking.ModuleName)

	// Sets the order of Genesis - Order matters, genutil is to always come last
	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	app.mm.SetOrderInitGenesis(
		capability.ModuleName, auth.ModuleName, distr.ModuleName, staking.ModuleName, bank.ModuleName,
		slashing.ModuleName, gov.ModuleName, crisis.ModuleName,
		ibc.ModuleName, genutil.ModuleName, evidence.ModuleName, transfer.ModuleName,
	)

	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.accountKeeper),
		bank.NewAppModule(appCodec, app.bankKeeper, app.accountKeeper),
		capability.NewAppModule(appCodec, *app.capabilityKeeper),
		gov.NewAppModule(appCodec, app.govKeeper, app.accountKeeper, app.bankKeeper),
		staking.NewAppModule(appCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		distr.NewAppModule(appCodec, app.distrKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		slashing.NewAppModule(appCodec, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		params.NewAppModule(app.paramsKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
	)

	app.sm.RegisterStoreDecoders()

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		ante.NewAnteHandler(
			app.accountKeeper, app.bankKeeper, *app.ibcKeeper,
			ante.DefaultSigVerificationGasConsumer,
		),
	)
	app.SetEndBlocker(app.EndBlocker)

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tKeys)
	app.MountMemoryStores(memKeys)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	ctx := app.BaseApp.NewUncachedContext(true, abci.Header{})
	app.capabilityKeeper.InitializeAndSeal(ctx)

	app.scopedIBCKeeper = scopedIBCKeeper
	app.scopedTransferKeeper = scopedTransferKeeper

	return app
}

// GenesisState represents chain state at the start of the chain. Any initial state (account balances) are stored here.
type GenesisState map[string]json.RawMessage

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState() GenesisState {
	cdc := std.MakeCodec(ModuleBasics)
	return ModuleBasics.DefaultGenesis(cdc)
}

// InitChainer application update at chain initialization
func (app *TsukiApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState simapp.GenesisState
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &genesisState)
	return app.mm.InitGenesis(ctx, app.cdc, genesisState)
}

// BeginBlocker application updates every begin block
func (app *TsukiApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block
func (app *TsukiApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// LoadHeight loads a particular height
func (app *TsukiApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *TsukiApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[auth.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// Codec returns the application's sealed codec.
func (app *TsukiApp) Codec() *codec.Codec {
	return app.cdc
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *TsukiApp) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.paramsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface
func (app *TsukiApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// GetMaccPerms returns a mapping of the application's module account permissions.
func GetMaccPerms() map[string][]string {
	modAccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		modAccPerms[k] = v
	}
	return modAccPerms
}

// BlacklistedAccAddrs returns all the app's module account addresses black listed for receiving tokens.
func (app *TsukiApp) BlacklistedAccAddrs() map[string]bool {
	blacklistedAddrs := make(map[string]bool)
	for acc := range maccPerms {
		blacklistedAddrs[auth.NewModuleAddress(acc).String()] = !allowedReceivingModAcc[acc]
	}

	return blacklistedAddrs
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.Marshaler, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)

	return paramsKeeper
}
