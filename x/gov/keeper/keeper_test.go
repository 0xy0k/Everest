package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	simapp "github.com/TsukiCore/tsuki/app"
	"github.com/TsukiCore/tsuki/x/gov/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

func TestKeeper_SetNetworkProperty(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	app.CustomGovKeeper.SetNetworkProperties(ctx, &types.NetworkProperties{
		MinTxFee: 100,
		MaxTxFee: 50000,
	})

	err := app.CustomGovKeeper.SetNetworkProperty(ctx, types.MinTxFee, types.NetworkPropertyValue{Value: 300})
	require.Nil(t, err)

	savedMinTxFee, err := app.CustomGovKeeper.GetNetworkProperty(ctx, types.MinTxFee)
	require.Nil(t, err)
	require.Equal(t, uint64(300), savedMinTxFee.Value)
}
