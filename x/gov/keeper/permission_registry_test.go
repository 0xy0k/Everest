package keeper_test

import (
	"testing"

	"github.com/TsukiCore/tsuki/simapp"
	"github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
)

func TestKeeper_SaveGetPermissionsForRole(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	perm := types.NewPermissions(
		nil, []types.PermValue{types.PermClaimValidator},
	)

	app.CustomGovKeeper.SetPermissionsForRole(ctx, types.RoleSudo, perm)

	savedPerms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.Equal(t, perm, &savedPerms)
}

func TestKeeper_HasGenesisDefaultRoles(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	roleSudo, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.True(t, roleSudo.IsWhitelisted(types.PermSetPermissions))

	roleValidator, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.True(t, roleValidator.IsWhitelisted(types.PermClaimValidator))
}

func TestKeeper_SetPermissionsOverwritesOldPerms(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	roleValidator, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.True(t, roleValidator.IsWhitelisted(types.PermClaimValidator))
	require.False(t, roleValidator.IsWhitelisted(types.PermSetPermissions))

	// We set whitelisted PermSetPermissions and Blacklisted PermClaimValidator
	newPerms := types.NewPermissions([]types.PermValue{types.PermSetPermissions}, []types.PermValue{types.PermClaimValidator})
	app.CustomGovKeeper.SetPermissionsForRole(ctx, types.RoleValidator, newPerms)

	newRoleValidatorPerms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.True(t, newRoleValidatorPerms.IsWhitelisted(types.PermSetPermissions))
	require.False(t, newRoleValidatorPerms.IsWhitelisted(types.PermClaimValidator))
}

func TestKeeper_GetPermissionsForRole_ReturnsNilWhenDoesNotExist(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	_, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, 12345)
	require.False(t, found)
}
