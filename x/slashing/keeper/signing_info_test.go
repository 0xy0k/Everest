package keeper_test

import (
	"testing"
	"time"

	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"

	simapp "github.com/TsukiCore/tsuki/app"
	"github.com/TsukiCore/tsuki/x/slashing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestGetSetValidatorSigningInfo(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	addrDels := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction))

	info, found := app.CustomSlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(addrDels[0]))
	require.False(t, found)
	newInfo := types.NewValidatorSigningInfo(
		sdk.ConsAddress(addrDels[0]),
		int64(4),
		time.Unix(2, 0),
		int64(10),
		int64(10),
		int64(10),
	)
	app.CustomSlashingKeeper.SetValidatorSigningInfo(ctx, sdk.ConsAddress(addrDels[0]), newInfo)
	info, found = app.CustomSlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(addrDels[0]))
	require.True(t, found)
	require.Equal(t, info.StartHeight, int64(4))
	require.Equal(t, info.InactiveUntil, time.Unix(2, 0).UTC())
	require.Equal(t, info.Mischance, int64(10))
	require.Equal(t, info.MissedBlocksCounter, int64(10))
	require.Equal(t, info.ProducedBlocksCounter, int64(10))
}

func TestJailUntil(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	addrDels := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction))

	require.Panics(t, func() { app.CustomSlashingKeeper.JailUntil(ctx, sdk.ConsAddress(addrDels[0]), time.Now()) })

	newInfo := types.NewValidatorSigningInfo(
		sdk.ConsAddress(addrDels[0]),
		int64(4),
		time.Unix(2, 0),
		int64(10),
		int64(10),
		int64(10),
	)
	app.CustomSlashingKeeper.SetValidatorSigningInfo(ctx, sdk.ConsAddress(addrDels[0]), newInfo)
	app.CustomSlashingKeeper.JailUntil(ctx, sdk.ConsAddress(addrDels[0]), time.Unix(253402300799, 0).UTC())

	info, ok := app.CustomSlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(addrDels[0]))
	require.True(t, ok)
	require.Equal(t, time.Unix(253402300799, 0).UTC(), info.InactiveUntil)
}
