package keeper_test

import (
	"testing"

	"github.com/TsukiCore/tsuki/simapp"
	"github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestKeeper_SaveCouncilor(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(10))
	addr := addrs[0]

	councilor := types.NewCouncilor(
		"moniker",
		addr,
	)

	app.CustomGovKeeper.SaveCouncilor(ctx, councilor)

	savedCouncilor, found := app.CustomGovKeeper.GetCouncilor(ctx, councilor.Address)
	require.True(t, found)
	require.Equal(t, councilor, savedCouncilor)

	// Get by moniker
	councilorByMoniker, found := app.CustomGovKeeper.GetCouncilorByMoniker(ctx, councilor.Moniker)
	require.True(t, found)
	require.Equal(t, councilor, councilorByMoniker)
}
