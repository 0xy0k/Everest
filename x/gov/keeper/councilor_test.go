package keeper

import (
	"testing"

	"github.com/TsukiCore/tsuki/simapp"
	"github.com/TsukiCore/tsuki/x/gov/types"
	types2 "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestKeeper_SaveCouncilor(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, types2.TokensFromConsensusPower(10))
	addr := addrs[0]

	councilor := types.NewCouncilor(
		"moniker",
		"website",
		"social",
		"identity",
		addr,
	)

	app.CustomGovKeeper.SaveCouncilor(ctx, councilor)

	savedCouncilor, err := app.CustomGovKeeper.GetCouncilor(ctx, councilor.Address)
	require.NoError(t, err)

	require.Equal(t, councilor, savedCouncilor)
}
