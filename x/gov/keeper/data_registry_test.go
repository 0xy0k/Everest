package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	simapp "github.com/TsukiCore/tsuki/app"
	"github.com/TsukiCore/tsuki/x/gov/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

func TestKeeper_UpsertDataRegistryEntry(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	entry := types.NewDataRegistryEntry(
		"someHAsh",
		"someURL",
		"someEncoding",
		1234,
	)

	app.CustomGovKeeper.UpsertDataRegistryEntry(ctx, "CodeOfConduct", entry)

	savedDataRegistry, found := app.CustomGovKeeper.GetDataRegistryEntry(ctx, "CodeOfConduct")
	require.True(t, found)

	require.Equal(t, entry, savedDataRegistry)

	_, found = app.CustomGovKeeper.GetDataRegistryEntry(ctx, "NonExistingKey")
	require.False(t, found)
}
