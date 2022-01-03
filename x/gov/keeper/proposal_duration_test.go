package keeper_test

import (
	"testing"

	simapp "github.com/TsukiCore/tsuki/app"
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestProposalDurationSetGet(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	// get not specifically define type
	duration := app.CustomGovKeeper.GetProposalDuration(ctx, tsukitypes.SetProposalDurationsProposalType)
	require.Equal(t, duration, uint64(0))

	// try to set correct value
	err := app.CustomGovKeeper.SetProposalDuration(ctx, tsukitypes.SetProposalDurationsProposalType, 2400)
	require.NoError(t, err)

	duration = app.CustomGovKeeper.GetProposalDuration(ctx, tsukitypes.SetProposalDurationsProposalType)
	require.Equal(t, duration, uint64(2400))

	// try setting again with lower than minimum value
	err = app.CustomGovKeeper.SetProposalDuration(ctx, tsukitypes.SetProposalDurationsProposalType, 1)
	require.Error(t, err)

	duration = app.CustomGovKeeper.GetProposalDuration(ctx, tsukitypes.SetProposalDurationsProposalType)
	require.Equal(t, duration, uint64(2400))

	// check get all functionality
	allDurations := app.CustomGovKeeper.GetAllProposalDurations(ctx)
	require.Equal(t, len(allDurations), 1)
}
