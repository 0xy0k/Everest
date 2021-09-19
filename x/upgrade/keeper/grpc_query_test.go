package keeper_test

import (
	"testing"

	simapp "github.com/TsukiCore/tsuki/app"
	"github.com/TsukiCore/tsuki/x/upgrade/keeper"
	"github.com/TsukiCore/tsuki/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestGRPCCurrentPlan(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	querier := keeper.NewQuerier(app.UpgradeKeeper)
	resp, err := querier.CurrentPlan(sdk.WrapSDKContext(ctx), &types.QueryCurrentPlanRequest{})
	require.NoError(t, err)
	require.Equal(t, resp, &types.QueryCurrentPlanResponse{Plan: nil})

	newPlan := types.Plan{
		UpgradeTime:          1,
		RollbackChecksum:     "checksum",
		MaxEnrolmentDuration: 2,
		Name:                 "plan",
	}

	app.UpgradeKeeper.SaveCurrentPlan(ctx, newPlan)

	resp, err = querier.CurrentPlan(sdk.WrapSDKContext(ctx), &types.QueryCurrentPlanRequest{})
	require.NoError(t, err)
	require.Equal(t, resp, &types.QueryCurrentPlanResponse{Plan: &newPlan})
}

func TestGRPCNextPlan(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	querier := keeper.NewQuerier(app.UpgradeKeeper)
	resp, err := querier.NextPlan(sdk.WrapSDKContext(ctx), &types.QueryNextPlanRequest{})
	require.NoError(t, err)
	require.Equal(t, resp, &types.QueryNextPlanResponse{Plan: nil})

	newPlan := types.Plan{
		UpgradeTime:          1,
		RollbackChecksum:     "checksum",
		MaxEnrolmentDuration: 2,
		Name:                 "plan",
	}

	app.UpgradeKeeper.SaveNextPlan(ctx, newPlan)

	resp, err = querier.NextPlan(sdk.WrapSDKContext(ctx), &types.QueryNextPlanRequest{})
	require.NoError(t, err)
	require.Equal(t, resp, &types.QueryNextPlanResponse{Plan: &newPlan})
}
