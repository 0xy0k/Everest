package genutil

import (
	"encoding/json"
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"

	"github.com/TsukiCore/tsuki/x/genutil/types"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	bankexported "github.com/cosmos/cosmos-sdk/x/bank/exported"
)

// SetGenTxsInAppGenesisState - sets the genesis transactions in the app genesis state
func SetGenTxsInAppGenesisState(
	cdc codec.JSONCodec, txJSONEncoder sdk.TxEncoder, appGenesisState map[string]json.RawMessage, genTxs []sdk.Tx,
) (map[string]json.RawMessage, error) {

	genesisState := types.GetGenesisStateFromAppState(cdc, appGenesisState)
	genTxsBz := make([]json.RawMessage, 0, len(genTxs))

	for _, genTx := range genTxs {
		txBz, err := txJSONEncoder(genTx)
		if err != nil {
			return appGenesisState, err
		}

		genTxsBz = append(genTxsBz, txBz)
	}

	genesisState.GenTxs = genTxsBz
	return types.SetGenesisStateInAppState(cdc, appGenesisState, genesisState), nil
}

// ValidateAccountInGenesis checks that the provided account has a sufficient
// balance in the set of genesis accounts.
func ValidateAccountInGenesis(
	appGenesisState map[string]json.RawMessage, genBalIterator types.GenesisBalancesIterator,
	addr sdk.Address, coins sdk.Coins, cdc codec.JSONCodec,
) error {
	var err error
	_, defaultDenom := govtypes.GetBech32PrefixAndDefaultDenomFromAppState(appGenesisState)

	accountIsInGenesis := false
	genBalIterator.IterateGenesisBalances(cdc, appGenesisState,
		func(bal bankexported.GenesisBalance) (stop bool) {
			accAddress := bal.GetAddress()
			accCoins := bal.GetCoins()

			// ensure that account is in genesis
			if accAddress.Equals(addr) {
				// ensure account contains enough funds of default bond denom
				if coins.AmountOf(defaultDenom).GT(accCoins.AmountOf(defaultDenom)) {
					err = fmt.Errorf(
						"account %s has a balance in genesis, but it only has %v%s available to stake, not %v%s",
						addr, accCoins.AmountOf(defaultDenom), defaultDenom, coins.AmountOf(defaultDenom), defaultDenom,
					)

					return true
				}

				accountIsInGenesis = true
				return true
			}

			return false
		},
	)

	if err != nil {
		return err
	}

	if !accountIsInGenesis {
		return fmt.Errorf("account %s does not have a balance in the genesis state", addr)
	}

	return nil
}

type deliverTxfn func(abci.RequestDeliverTx) abci.ResponseDeliverTx

// DeliverGenTxs iterates over all genesis txs, decodes each into a Tx and
// invokes the provided deliverTxfn with the decoded Tx. It returns the result
// of the staking module's ApplyAndReturnValidatorSetUpdates.
func DeliverGenTxs(
	ctx sdk.Context, genTxs []json.RawMessage,
	stakingKeeper types.StakingKeeper, deliverTx deliverTxfn,
	txEncodingConfig client.TxEncodingConfig,
) ([]abci.ValidatorUpdate, error) {

	for _, genTx := range genTxs {
		tx, err := txEncodingConfig.TxJSONDecoder()(genTx)
		if err != nil {
			panic(err)
		}

		bz, err := txEncodingConfig.TxEncoder()(tx)
		if err != nil {
			panic(err)
		}

		res := deliverTx(abci.RequestDeliverTx{Tx: bz})
		if !res.IsOK() {
			panic(res.Log)
		}
	}

	return stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
}
