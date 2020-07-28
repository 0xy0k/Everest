package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/cast"

	banktypes "github.com/TsukiCore/cosmos-sdk/x/bank/types"

	"github.com/TsukiCore/cosmos-sdk/x/auth/types"

	"github.com/TsukiCore/cosmos-sdk/client"
	"github.com/TsukiCore/cosmos-sdk/simapp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/TsukiCore/tsuki/app"

	"github.com/TsukiCore/cosmos-sdk/baseapp"
	"github.com/TsukiCore/cosmos-sdk/client/debug"
	"github.com/TsukiCore/cosmos-sdk/client/flags"
	"github.com/TsukiCore/cosmos-sdk/server"
	"github.com/TsukiCore/cosmos-sdk/store"
	sdk "github.com/TsukiCore/cosmos-sdk/types"
	genutilcli "github.com/TsukiCore/cosmos-sdk/x/genutil/client/cli"
)

const flagInvCheckPeriod = "inv-check-period"

var (
	invCheckPeriod uint
	rootCmd        = &cobra.Command{
		Use:   "tsukid",
		Short: "Tsuki Daemon (server)",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			return server.InterceptConfigsPreRunHandler(cmd)
		},
	}

	encodingConfig = simapp.MakeEncodingConfig()
	initClientCtx  = client.Context{}.
			WithJSONMarshaler(encodingConfig.Marshaler).
			WithTxConfig(encodingConfig.TxConfig).
			WithCodec(encodingConfig.Amino).
			WithInput(os.Stdin).
			WithAccountRetriever(types.NewAccountRetriever(encodingConfig.Marshaler)).
			WithBroadcastMode(flags.BroadcastBlock).
			WithHomeDir(simapp.DefaultNodeHome)
)

func main() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	rootCmd.AddCommand(
		genutilcli.InitCmd(app.ModuleBasics, app.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.MigrateGenesisCmd(),
		genutilcli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, app.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(app.ModuleBasics, encodingConfig.TxConfig),
		AddGenesisAccountCmd(app.DefaultNodeHome),
		cli.NewCompletionCmd(rootCmd, true),
		testnetCmd(app.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		replayCmd(),
		debug.Cmd(),
	)

	server.AddCommands(rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "GA", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer, appOpts server.AppOptions) server.Application {
	var cache sdk.MultiStorePersistentCache

	if viper.GetBool(server.FlagInterBlockCache) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range viper.GetIntSlice(server.FlagUnsafeSkipUpgrades) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	return app.NewInitApp(
		logger, db, traceStore, true, invCheckPeriod, skipUpgradeHeights,
		viper.GetString(flags.FlagHome),
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(viper.GetUint64(server.FlagHaltHeight)),
		baseapp.SetHaltTime(viper.GetUint64(server.FlagHaltTime)),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, *abci.ConsensusParams, error) {

	if height != -1 {
		tsukiapp := app.NewInitApp(logger, db, traceStore, false, uint(1), map[int64]bool{}, "")
		err := tsukiapp.LoadHeight(height)
		if err != nil {
			return nil, nil, nil, err
		}

		return tsukiapp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	tsukiapp := app.NewInitApp(logger, db, traceStore, true, uint(1), map[int64]bool{}, "")
	return tsukiapp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
