package types_test

import (
	"testing"

	stakingtypes "github.com/TsukiCore/tsuki/x/staking/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewValidator_IsActiveByDefault(t *testing.T) {
	valAddr, err := types.ValAddressFromBech32("tsukivaloper1q24436yrnettd6v4eu6r4t9gycnnddac9nwqv0")
	require.NoError(t, err)

	pubkeys := simtestutil.CreateTestPubKeys(1)
	pubKey := pubkeys[0]

	validator, err := stakingtypes.NewValidator(
		valAddr,
		pubKey,
	)
	require.NoError(t, err)
	require.True(t, validator.IsActive())
}
