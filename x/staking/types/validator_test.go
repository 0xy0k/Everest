package types_test

import (
	"testing"

	stakingtypes "github.com/TsukiCore/tsuki/x/staking/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestNewValidator_IsActiveByDefault(t *testing.T) {
	valAddr, err := types.ValAddressFromBech32("tsukivaloper1q24436yrnettd6v4eu6r4t9gycnnddac9nwqv0")
	require.NoError(t, err)

	pubkeys := simapp.CreateTestPubKeys(1)
	pubKey := pubkeys[0]

	validator, err := stakingtypes.NewValidator(
		valAddr,
		pubKey,
		sdk.NewDecWithPrec(5, 2),
	)
	require.NoError(t, err)
	require.True(t, validator.IsActive())
}
