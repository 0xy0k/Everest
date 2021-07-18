package keeper_test

import (
	"testing"

	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	stakingkeeper "github.com/TsukiCore/tsuki/x/staking/keeper"

	stakingtypes "github.com/TsukiCore/tsuki/x/staking/types"

	"github.com/TsukiCore/tsuki/simapp"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestQuerier_ValidatorByAddress(t *testing.T) {
	valAddr1, err := types.ValAddressFromBech32("tsukivaloper15ky9du8a2wlstz6fpx3p4mqpjyrm5cgq38f2fp")
	require.NoError(t, err)
	pubKey, err := types.GetPubKeyFromBech32(types.Bech32PubKeyTypeConsPub, "tsukivalconspub1zcjduepqylc5k8r40azmw0xt7hjugr4mr5w2am7jw77ux5w6s8hpjxyrjjsq4xg7em")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	val, err := stakingtypes.NewValidator("Moniker", types.NewDec(123), valAddr1, pubKey)
	require.NoError(t, err)

	app.CustomStakingKeeper.AddValidator(ctx, val)

	querier := stakingkeeper.NewQuerier(app.CustomStakingKeeper)

	qValidatorResp, err := querier.ValidatorByAddress(types.WrapSDKContext(ctx), &stakingtypes.ValidatorByAddressRequest{ValAddr: valAddr1})
	require.NoError(t, err)

	require.True(t, val.Equal(qValidatorResp.Validator))
}
