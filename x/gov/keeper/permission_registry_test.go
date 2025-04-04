package keeper_test

import (
	"testing"

	simapp "github.com/TsukiCore/tsuki/app"
	"github.com/TsukiCore/tsuki/x/gov/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_CreateRoleAndWhitelistPerm(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	app.CustomGovKeeper.SetRole(ctx, types.Role{
		Id:          uint32(types.RoleSudo),
		Sid:         "RoleSudo",
		Description: "RoleSudo",
	})

	err := app.CustomGovKeeper.WhitelistRolePermission(ctx, types.RoleSudo, types.PermClaimValidator)
	require.NoError(t, err)

	savedPerms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.True(t, savedPerms.IsWhitelisted(types.PermClaimValidator))
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

	iterator := app.CustomGovKeeper.GetRolesByWhitelistedPerm(ctx, types.PermClaimValidator)
	requireIteratorCount(t, iterator, 2)
}

func TestKeeper_WhitelistRolePermission(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	perms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.False(t, perms.IsWhitelisted(types.PermChangeTxFee))

	err := app.CustomGovKeeper.WhitelistRolePermission(ctx, types.RoleSudo, types.PermChangeTxFee)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.True(t, perms.IsWhitelisted(types.PermChangeTxFee))

	iterator := app.CustomGovKeeper.GetRolesByWhitelistedPerm(ctx, types.PermChangeTxFee)
	requireIteratorCount(t, iterator, 1)
}

func TestKeeper_RemoveWhitelistRolePermission(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	perms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.False(t, perms.IsWhitelisted(types.PermChangeTxFee))

	err := app.CustomGovKeeper.WhitelistRolePermission(ctx, types.RoleSudo, types.PermChangeTxFee)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.True(t, perms.IsWhitelisted(types.PermChangeTxFee))

	iterator := app.CustomGovKeeper.GetRolesByWhitelistedPerm(ctx, types.PermChangeTxFee)
	requireIteratorCount(t, iterator, 1)

	err = app.CustomGovKeeper.RemoveWhitelistRolePermission(ctx, types.RoleSudo, types.PermChangeTxFee)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.False(t, perms.IsWhitelisted(types.PermChangeTxFee))

	iterator = app.CustomGovKeeper.GetRolesByWhitelistedPerm(ctx, types.PermChangeTxFee)
	requireIteratorCount(t, iterator, 0)
}

func TestKeeper_BlacklistRolePermission(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	perms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.False(t, perms.IsBlacklisted(types.PermChangeTxFee))

	err := app.CustomGovKeeper.BlacklistRolePermission(ctx, types.RoleSudo, types.PermChangeTxFee)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.True(t, perms.IsBlacklisted(types.PermChangeTxFee))
}

func TestKeeper_RemoveBlacklistRolePermission(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	perms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.False(t, perms.IsBlacklisted(types.PermChangeTxFee))

	err := app.CustomGovKeeper.BlacklistRolePermission(ctx, types.RoleSudo, types.PermChangeTxFee)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.True(t, perms.IsBlacklisted(types.PermChangeTxFee))

	err = app.CustomGovKeeper.RemoveBlacklistRolePermission(ctx, types.RoleSudo, types.PermChangeTxFee)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleSudo)
	require.True(t, found)
	require.False(t, perms.IsBlacklisted(types.PermChangeTxFee))
}

func TestKeeper_GetPermissionsForRole_ReturnsNilWhenDoesNotExist(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	_, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, 12345)
	require.False(t, found)
}
