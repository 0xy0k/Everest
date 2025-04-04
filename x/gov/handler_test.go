package gov_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	simapp "github.com/TsukiCore/tsuki/app"
	appparams "github.com/TsukiCore/tsuki/app/params"
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov"
	"github.com/TsukiCore/tsuki/x/gov/types"
	tokenstypes "github.com/TsukiCore/tsuki/x/tokens/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	appparams.SetConfig()
	os.Exit(m.Run())
}

// When a network actor has not been saved before, it creates one with default params
// and sets the permissions.
func TestHandler_MsgWhitelistPermissions_ActorDoesNotExist(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name             string
		msg              sdk.Msg
		checkWhitelisted bool
	}{
		{
			"Msg Whitelist Permissions",
			&types.MsgWhitelistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			true,
		},
		{
			"Msg Blacklist Permissions",
			&types.MsgBlacklistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			err := setPermissionToAddr(t, app, ctx, proposerAddr, types.PermSetPermissions)
			require.NoError(t, err)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err = handler(ctx, tt.msg)
			require.NoError(t, err)

			actor, found := app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
			require.True(t, found)

			if tt.checkWhitelisted {
				require.True(t, actor.Permissions.IsWhitelisted(types.PermClaimValidator))
			} else {
				require.True(t, actor.Permissions.IsBlacklisted(types.PermClaimValidator))
			}
		})
	}
}

// When a network actor has already permissions it just appends the permission.
func TestNewHandler_SetPermissions_ActorWithPerms(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name             string
		msg              sdk.Msg
		checkWhitelisted bool
	}{
		{
			name: "actor with Whitelist Permissions",
			msg: &types.MsgWhitelistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			checkWhitelisted: true,
		},
		{
			name: "actor with Blacklist Permissions",
			msg: &types.MsgBlacklistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			checkWhitelisted: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			err := setPermissionToAddr(t, app, ctx, proposerAddr, types.PermSetPermissions)
			require.NoError(t, err)

			// Add some perms before to the actor.
			actor := types.NewDefaultActor(addr)
			if tt.checkWhitelisted {
				err = actor.Permissions.AddToWhitelist(types.PermSetPermissions)
			} else {
				err = actor.Permissions.AddToBlacklist(types.PermSetPermissions)
			}
			require.NoError(t, err)

			app.CustomGovKeeper.SaveNetworkActor(ctx, actor)

			// Call the handler to add some permissions.
			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err = handler(ctx, tt.msg)
			require.NoError(t, err)

			actor, found := app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
			require.True(t, found)

			if tt.checkWhitelisted {
				require.True(t, actor.Permissions.IsWhitelisted(types.PermClaimValidator))
				require.True(t, actor.Permissions.IsWhitelisted(types.PermSetPermissions)) // This permission was already set before callid add permission.
			} else {
				require.True(t, actor.Permissions.IsBlacklisted(types.PermClaimValidator))
				require.True(t, actor.Permissions.IsBlacklisted(types.PermSetPermissions)) // This permission was already set before callid add permission.
			}
		})
	}
}

func TestNewHandler_SetPermissionsWithoutSetPermissions(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  sdk.Msg
	}{
		{
			name: "MsgWhitelist",
			msg: &types.MsgWhitelistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
		},
		{
			name: "MsgBlacklist",
			msg: &types.MsgBlacklistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err = handler(ctx, tt.msg)
			require.EqualError(t, err, "PermSetPermissions || (ClaimValidatorPermission && ClaimValidatorPermMsg): not enough permissions")
		})
	}
}

func TestNewHandler_SetPermissions_ProposerHasRoleSudo(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name           string
		msg            sdk.Msg
		checkWhitelist bool
	}{
		{
			name: "MsgWhitelist",
			msg: &types.MsgWhitelistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			checkWhitelist: true,
		},
		{
			name: "MsgBlacklist",
			msg: &types.MsgBlacklistPermissions{
				Proposer:   proposerAddr,
				Address:    addr,
				Permission: uint32(types.PermClaimValidator),
			},
			checkWhitelist: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			// First we set Role Sudo to proposer Actor
			proposerActor := types.NewDefaultActor(proposerAddr)
			proposerActor.SetRole(types.RoleSudo)
			require.NoError(t, err)
			app.CustomGovKeeper.SaveNetworkActor(ctx, proposerActor)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err = handler(ctx, tt.msg)
			require.NoError(t, err)

			actor, found := app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
			require.True(t, found)

			if tt.checkWhitelist {
				require.True(t, actor.Permissions.IsWhitelisted(types.PermClaimValidator))
			} else {
				require.True(t, actor.Permissions.IsBlacklisted(types.PermClaimValidator))
			}
		})
	}
}

func TestNewHandler_SetNetworkProperties(t *testing.T) {
	changeFeeAddr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	sudoAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name       string
		msg        sdk.Msg
		desiredErr string
	}{
		{
			name: "Success run with ChangeTxFee permission",
			msg: &types.MsgSetNetworkProperties{
				NetworkProperties: &types.NetworkProperties{
					MinTxFee: 100,
					MaxTxFee: 1000,
				},
				Proposer: changeFeeAddr,
			},
			desiredErr: "",
		},
		{
			name: "Failure run without ChangeTxFee permission",
			msg: &types.MsgSetNetworkProperties{
				NetworkProperties: &types.NetworkProperties{
					MinTxFee: 100,
					MaxTxFee: 1000,
				},
				Proposer: sudoAddr,
			},
			desiredErr: "not enough permissions",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})
			// First we set Role Sudo to proposer Actor
			proposerActor := types.NewDefaultActor(sudoAddr)
			proposerActor.SetRole(types.RoleSudo)
			app.CustomGovKeeper.SaveNetworkActor(ctx, proposerActor)

			handler := gov.NewHandler(app.CustomGovKeeper)

			// set change fee permission to addr
			_, err = handler(ctx, &types.MsgWhitelistPermissions{
				Proposer:   sudoAddr,
				Address:    changeFeeAddr,
				Permission: uint32(types.PermChangeTxFee),
			})
			require.NoError(t, err)

			_, err = handler(ctx, tt.msg)
			if tt.desiredErr == "" {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.desiredErr)
			}
		})
	}
}

func TestNewHandler_SetExecutionFee(t *testing.T) {
	execFeeSetAddr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	sudoAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name       string
		msg        types.MsgSetExecutionFee
		desiredErr string
	}{
		{
			name: "Success run with ChangeTxFee permission",
			msg: types.MsgSetExecutionFee{
				TransactionType:   "network-properties",
				ExecutionFee:      10000,
				FailureFee:        1000,
				Timeout:           1,
				DefaultParameters: 2,
				Proposer:          execFeeSetAddr,
			},
			desiredErr: "",
		},
		{
			name: "Success run without ChangeTxFee permission",
			msg: types.MsgSetExecutionFee{
				TransactionType:   "network-properties-2",
				ExecutionFee:      10000,
				FailureFee:        1000,
				Timeout:           1,
				DefaultParameters: 2,
				Proposer:          sudoAddr,
			},
			desiredErr: "PermChangeTxFee: not enough permissions",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})
			// First we set Role Sudo to proposer Actor
			proposerActor := types.NewDefaultActor(sudoAddr)
			proposerActor.SetRole(types.RoleSudo)
			require.NoError(t, err)
			app.CustomGovKeeper.SaveNetworkActor(ctx, proposerActor)

			handler := gov.NewHandler(app.CustomGovKeeper)

			// set change fee permission to addr
			_, err = handler(ctx, &types.MsgWhitelistPermissions{
				Proposer:   sudoAddr,
				Address:    execFeeSetAddr,
				Permission: uint32(types.PermChangeTxFee),
			})
			require.NoError(t, err)

			_, err = handler(ctx, &tt.msg)
			if tt.desiredErr == "" {
				require.NoError(t, err)
				execFee := app.CustomGovKeeper.GetExecutionFee(ctx, tt.msg.TransactionType)
				require.Equal(t, tt.msg.TransactionType, execFee.TransactionType)
				require.Equal(t, tt.msg.ExecutionFee, execFee.ExecutionFee)
				require.Equal(t, tt.msg.FailureFee, execFee.FailureFee)
				require.Equal(t, tt.msg.Timeout, execFee.Timeout)
				require.Equal(t, tt.msg.DefaultParameters, execFee.DefaultParameters)
			} else {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.desiredErr)
			}
		})
	}
}

func TestHandler_ClaimCouncilor_Fails(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name        string
		msg         sdk.Msg
		expectedErr error
	}{
		{
			name: "not enough permissions",
			msg: &types.MsgClaimCouncilor{
				Moniker: "",
				Address: addr,
			},
			expectedErr: fmt.Errorf("PermClaimCouncilor: not enough permissions"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err := handler(ctx, tt.msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_ClaimCouncilor_HappyPath(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  *types.MsgClaimCouncilor
	}{
		{
			name: "all correct",
			msg: &types.MsgClaimCouncilor{
				Address: addr,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			err = setPermissionToAddr(t, app, ctx, addr, types.PermClaimCouncilor)
			require.NoError(t, err)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err := handler(ctx, tt.msg)
			require.NoError(t, err)

			expectedCouncilor := types.NewCouncilor(
				tt.msg.Address,
				types.CouncilorActive,
			)

			councilor, found := app.CustomGovKeeper.GetCouncilor(ctx, addr)
			require.True(t, found)

			require.Equal(t, expectedCouncilor, councilor)
		})
	}
}

func TestHandler_WhitelistRolePermissions_Errors(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		msg          *types.MsgWhitelistRolePermission
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			name: "address without SetPermissions perm",
			msg: types.NewMsgWhitelistRolePermission(
				addr,
				fmt.Sprintf("%d", types.RoleValidator),
				uint32(types.PermSetPermissions),
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				return
			},
			expectedErr: fmt.Errorf("%s: not enough permissions", types.PermUpsertRole.String()),
		},
		{
			name: "role does not exist",
			msg: types.NewMsgWhitelistRolePermission(
				addr,
				"10000",
				1,
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
				require.NoError(t, err2)
			},
			expectedErr: fmt.Errorf("role does not exist"),
		},
		{
			name: "permission is blacklisted",
			msg: types.NewMsgWhitelistRolePermission(
				addr,
				fmt.Sprintf("%d", types.RoleValidator),
				uint32(types.PermSetPermissions),
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
				require.NoError(t, err2)

				_, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
				require.True(t, found)

				err2 = app.CustomGovKeeper.BlacklistRolePermission(ctx, types.RoleValidator, types.PermSetPermissions)
				require.NoError(t, err2)
			},
			expectedErr: fmt.Errorf("permission is blacklisted"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			tt.preparePerms(t, app, ctx)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err := handler(ctx, tt.msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_WhitelistRolePermissions(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	err = setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
	require.NoError(t, err)

	perms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.False(t, perms.IsWhitelisted(types.PermSetPermissions))

	msg := types.NewMsgWhitelistRolePermission(
		addr,
		fmt.Sprintf("%d", types.RoleValidator),
		uint32(types.PermSetPermissions),
	)

	handler := gov.NewHandler(app.CustomGovKeeper)
	_, err = handler(ctx, msg)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.True(t, perms.IsWhitelisted(types.PermSetPermissions))
}

func TestHandler_BlacklistRolePermissions_Errors(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		msg          *types.MsgBlacklistRolePermission
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			name: "address without SetPermissions perm",
			msg: types.NewMsgBlacklistRolePermission(
				addr,
				fmt.Sprintf("%d", types.RoleValidator),
				uint32(types.PermSetPermissions),
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			expectedErr:  fmt.Errorf("%s: not enough permissions", types.PermUpsertRole.String()),
		},
		{
			name: "role does not exist",
			msg: types.NewMsgBlacklistRolePermission(
				addr,
				"10000",
				1,
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
				require.NoError(t, err2)
			},
			expectedErr: fmt.Errorf("role does not exist"),
		},
		{
			name: "permission is whitelisted",
			msg: types.NewMsgBlacklistRolePermission(
				addr,
				fmt.Sprintf("%d", types.RoleValidator),
				uint32(types.PermSetPermissions),
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
				require.NoError(t, err2)

				_, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
				require.True(t, found)

				err2 = app.CustomGovKeeper.WhitelistRolePermission(ctx, types.RoleValidator, types.PermSetPermissions)
				require.NoError(t, err2)
			},
			expectedErr: fmt.Errorf("permission is whitelisted"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			tt.preparePerms(t, app, ctx)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err := handler(ctx, tt.msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_BlacklistRolePermissions(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	err = setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
	require.NoError(t, err)

	perms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.False(t, perms.IsBlacklisted(types.PermSetPermissions))

	msg := types.NewMsgBlacklistRolePermission(
		addr,
		fmt.Sprintf("%d", types.RoleValidator),
		uint32(types.PermSetPermissions),
	)

	handler := gov.NewHandler(app.CustomGovKeeper)
	_, err = handler(ctx, msg)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.True(t, perms.IsBlacklisted(types.PermSetPermissions))
}

func TestHandler_RemoveWhitelistRolePermissions_Errors(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		msg          *types.MsgRemoveWhitelistRolePermission
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			name: "address without SetPermissions perm",
			msg: types.NewMsgRemoveWhitelistRolePermission(
				addr,
				fmt.Sprintf("%d", types.RoleValidator),
				uint32(types.PermSetPermissions),
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			expectedErr:  fmt.Errorf("%s: not enough permissions", types.PermUpsertRole.String()),
		},
		{
			name: "role does not exist",
			msg: types.NewMsgRemoveWhitelistRolePermission(
				addr,
				"10000",
				1,
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
				require.NoError(t, err2)
			},
			expectedErr: fmt.Errorf("role does not exist"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			tt.preparePerms(t, app, ctx)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err := handler(ctx, tt.msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_RemoveWhitelistRolePermissions(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	err = setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
	require.NoError(t, err)

	perms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.True(t, perms.IsWhitelisted(types.PermClaimValidator))

	msg := types.NewMsgRemoveWhitelistRolePermission(
		addr,
		fmt.Sprintf("%d", types.RoleValidator),
		uint32(types.PermClaimValidator),
	)

	handler := gov.NewHandler(app.CustomGovKeeper)
	_, err = handler(ctx, msg)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.False(t, perms.IsWhitelisted(types.PermClaimValidator))
}

func TestHandler_RemoveBlacklistRolePermissions_Errors(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		msg          *types.MsgRemoveBlacklistRolePermission
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			name: "address without SetPermissions perm",
			msg: types.NewMsgRemoveBlacklistRolePermission(
				addr,
				fmt.Sprintf("%d", types.RoleValidator),
				uint32(types.PermSetPermissions),
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			expectedErr:  fmt.Errorf("%s: not enough permissions", types.PermUpsertRole.String()),
		},
		{
			name: "role does not exist",
			msg: types.NewMsgRemoveBlacklistRolePermission(
				addr,
				"10000",
				1,
			),
			preparePerms: func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
				require.NoError(t, err2)
			},
			expectedErr: fmt.Errorf("role does not exist"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			tt.preparePerms(t, app, ctx)

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err := handler(ctx, tt.msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_RemoveBlacklistRolePermissions(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	err = setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
	require.NoError(t, err)

	_, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)

	// Set some blacklist value
	err = app.CustomGovKeeper.BlacklistRolePermission(ctx, types.RoleValidator, types.PermClaimCouncilor)
	require.NoError(t, err)

	// Check if it is blacklisted.
	perms, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.True(t, perms.IsBlacklisted(types.PermClaimCouncilor))

	msg := types.NewMsgRemoveBlacklistRolePermission(
		addr,
		fmt.Sprintf("%d", types.RoleValidator),
		uint32(types.PermClaimCouncilor),
	)

	handler := gov.NewHandler(app.CustomGovKeeper)
	_, err = handler(ctx, msg)
	require.NoError(t, err)

	perms, found = app.CustomGovKeeper.GetPermissionsForRole(ctx, types.RoleValidator)
	require.True(t, found)
	require.False(t, perms.IsBlacklisted(types.PermClaimCouncilor))
}

func TestHandler_CreateRole_Errors(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		msg          *types.MsgCreateRole
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			"fails when no perms",
			types.NewMsgCreateRole(
				addr,
				"role10",
				"role10desc",
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			fmt.Errorf("PermUpsertRole: not enough permissions"),
		},
		{
			"fails when role already exists",
			types.NewMsgCreateRole(
				addr,
				"role1234",
				"role1234desc",
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
				require.NoError(t, err2)
				app.CustomGovKeeper.CreateRole(ctx, "role1234", "role1234desc")
			},
			fmt.Errorf("role already exist"),
		},
	}

	for _, tt := range tests {
		app := simapp.Setup(false)
		ctx := app.NewContext(false, tmproto.Header{})

		tt.preparePerms(t, app, ctx)

		handler := gov.NewHandler(app.CustomGovKeeper)
		_, err := handler(ctx, tt.msg)
		require.EqualError(t, err, tt.expectedErr.Error())
	}
}

func TestHandler_CreateRole(t *testing.T) {
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	err = setPermissionToAddr(t, app, ctx, addr, types.PermUpsertRole)
	require.NoError(t, err)

	_, found := app.CustomGovKeeper.GetPermissionsForRole(ctx, 1234)
	require.False(t, found)

	handler := gov.NewHandler(app.CustomGovKeeper)
	_, err = handler(ctx, types.NewMsgCreateRole(
		addr,
		"role1234",
		"role1234desc",
	))
	require.NoError(t, err)

	_, err = app.CustomGovKeeper.GetRoleBySid(ctx, "role1234")
	require.NoError(t, err)
}

func TestHandler_AssignRole_Errors(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		msg          *types.MsgAssignRole
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			"fails when no perms",
			types.NewMsgAssignRole(
				proposerAddr, addr, 3,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			fmt.Errorf("%s: not enough permissions", types.PermUpsertRole.String()),
		},
		{
			"fails when role does not exist",
			types.NewMsgAssignRole(
				proposerAddr, addr, 3,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, proposerAddr, types.PermUpsertRole)
				require.NoError(t, err2)
			},
			types.ErrRoleDoesNotExist,
		},
		{
			"role already assigned",
			types.NewMsgAssignRole(
				proposerAddr, addr, 3,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, proposerAddr, types.PermUpsertRole)
				require.NoError(t, err2)

				app.CustomGovKeeper.SetRole(ctx, types.Role{
					Id:          3,
					Sid:         "3",
					Description: "3",
				})
				err2 = app.CustomGovKeeper.WhitelistRolePermission(ctx, 3, types.PermClaimValidator)
				require.NoError(t, err2)

				networkActor := types.NewDefaultActor(addr)
				app.CustomGovKeeper.AssignRoleToActor(ctx, networkActor, 3)
			},
			types.ErrRoleAlreadyAssigned,
		},
	}

	for _, tt := range tests {
		app := simapp.Setup(false)
		ctx := app.NewContext(false, tmproto.Header{})

		tt.preparePerms(t, app, ctx)

		handler := gov.NewHandler(app.CustomGovKeeper)
		_, err := handler(ctx, tt.msg)
		require.EqualError(t, err, tt.expectedErr.Error())
	}
}

func TestHandler_AssignRole(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	// Set permissions to proposer.
	err = setPermissionToAddr(t, app, ctx, proposerAddr, types.PermUpsertRole)
	require.NoError(t, err)

	// Create role
	app.CustomGovKeeper.SetRole(ctx, types.Role{
		Id:          3,
		Sid:         "3",
		Description: "3",
	})
	err = app.CustomGovKeeper.WhitelistRolePermission(ctx, 3, types.PermSetPermissions)
	require.NoError(t, err)

	msg := types.NewMsgAssignRole(proposerAddr, addr, 3)

	handler := gov.NewHandler(app.CustomGovKeeper)
	_, err = handler(ctx, msg)
	require.NoError(t, err)

	actor, found := app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
	require.True(t, found)

	require.True(t, actor.HasRole(3))
}

func TestHandler_UnassignRole_Errors(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		msg          *types.MsgUnassignRole
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			"fails when no perms",
			types.NewMsgUnassignRole(
				proposerAddr, addr, 3,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			fmt.Errorf("%s: not enough permissions", types.PermUpsertRole.String()),
		},
		{
			"fails when role does not exist",
			types.NewMsgUnassignRole(
				proposerAddr, addr, 3,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, proposerAddr, types.PermUpsertRole)
				require.NoError(t, err2)
			},
			types.ErrRoleDoesNotExist,
		},
		{
			"role not assigned",
			types.NewMsgUnassignRole(
				proposerAddr, addr, 3,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				err2 := setPermissionToAddr(t, app, ctx, proposerAddr, types.PermUpsertRole)
				require.NoError(t, err2)

				app.CustomGovKeeper.SetRole(ctx, types.Role{
					Id:          3,
					Sid:         "3",
					Description: "3",
				})
				err2 = app.CustomGovKeeper.WhitelistRolePermission(ctx, 3, types.PermClaimValidator)
				require.NoError(t, err2)
				networkActor := types.NewDefaultActor(addr)
				app.CustomGovKeeper.SaveNetworkActor(ctx, networkActor)
			},
			types.ErrRoleNotAssigned,
		},
	}

	for _, tt := range tests {
		app := simapp.Setup(false)
		ctx := app.NewContext(false, tmproto.Header{})

		tt.preparePerms(t, app, ctx)

		handler := gov.NewHandler(app.CustomGovKeeper)
		_, err := handler(ctx, tt.msg)
		require.EqualError(t, err, tt.expectedErr.Error())
	}
}

func TestHandler_UnassignRoles(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	// Set permissions to proposer.
	err = setPermissionToAddr(t, app, ctx, proposerAddr, types.PermUpsertRole)
	require.NoError(t, err)

	// Set new role and set permission to actor.
	app.CustomGovKeeper.SetRole(ctx, types.Role{
		Id:          3,
		Sid:         "3",
		Description: "3",
	})
	err = app.CustomGovKeeper.WhitelistRolePermission(ctx, 3, types.PermSetPermissions)
	require.NoError(t, err)

	actor := types.NewDefaultActor(addr)
	app.CustomGovKeeper.AssignRoleToActor(ctx, actor, 3)

	actor, found := app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
	require.True(t, found)
	require.True(t, actor.HasRole(3))

	msg := types.NewMsgUnassignRole(proposerAddr, addr, 3)
	handler := gov.NewHandler(app.CustomGovKeeper)
	_, err = handler(ctx, msg)
	require.NoError(t, err)

	actor, found = app.CustomGovKeeper.GetNetworkActorByAddress(ctx, addr)
	require.True(t, found)

	require.False(t, actor.HasRole(3))
}

func TestHandler_CreateProposalAssignPermission_Errors(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		content      types.Content
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			"Proposer does not have Perm",
			types.NewWhitelistAccountPermissionProposal(
				addr, types.PermClaimValidator,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			errors.Wrap(types.ErrNotEnoughPermissions, types.PermWhitelistAccountPermissionProposal.String()),
		},
		{
			"address already has that permission",
			types.NewWhitelistAccountPermissionProposal(
				addr, types.PermClaimValidator,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				proposerActor := types.NewDefaultActor(proposerAddr)
				err2 := app.CustomGovKeeper.AddWhitelistPermission(ctx, proposerActor, types.PermWhitelistAccountPermissionProposal)
				require.NoError(t, err2)

				actor := types.NewDefaultActor(addr)
				err2 = app.CustomGovKeeper.AddWhitelistPermission(ctx, actor, types.PermClaimValidator)
				require.NoError(t, err2)
			},
			fmt.Errorf("permission already whitelisted: error adding to whitelist"),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			tt.preparePerms(t, app, ctx)

			handler := gov.NewHandler(app.CustomGovKeeper)
			msg, err := types.NewMsgSubmitProposal(proposerAddr, "title", "some desc", tt.content)
			require.NoError(t, err)
			_, err = handler(ctx, msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_ProposalAssignPermission(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{
		Time: time.Now(),
	})

	// Set proposer Permissions
	proposerActor := types.NewDefaultActor(proposerAddr)
	err2 := app.CustomGovKeeper.AddWhitelistPermission(ctx, proposerActor, types.PermWhitelistAccountPermissionProposal)
	require.NoError(t, err2)

	properties := app.CustomGovKeeper.GetNetworkProperties(ctx)
	properties.MinimumProposalEndTime = 600 // Seconds, 10 mins
	app.CustomGovKeeper.SetNetworkProperties(ctx, properties)

	handler := gov.NewHandler(app.CustomGovKeeper)
	proposal := types.NewWhitelistAccountPermissionProposal(addr, types.PermValue(1))
	msg, err := types.NewMsgSubmitProposal(proposerAddr, "title", "some desc", proposal)
	require.NoError(t, err)
	res, err := handler(
		ctx,
		msg,
	)
	require.NoError(t, err)

	expData, _ := proto.Marshal(&types.MsgSubmitProposalResponse{ProposalID: 1})
	require.Equal(t, expData, res.Data)

	savedProposal, found := app.CustomGovKeeper.GetProposal(ctx, 1)
	require.True(t, found)

	expectedSavedProposal, err := types.NewProposal(
		1,
		"title",
		"some desc",
		types.NewWhitelistAccountPermissionProposal(
			addr,
			types.PermValue(1),
		),
		ctx.BlockTime(),
		ctx.BlockTime().Add(time.Second*time.Duration(properties.MinimumProposalEndTime)),
		ctx.BlockTime().Add((time.Second*time.Duration(properties.MinimumProposalEndTime))+(time.Second*time.Duration(properties.ProposalEnactmentTime))),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)
	require.Equal(t, expectedSavedProposal, savedProposal)

	// Next proposal ID is increased.
	id := app.CustomGovKeeper.GetNextProposalID(ctx)
	require.Equal(t, uint64(2), id)

	// Is not on finished active proposals.
	iterator := app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, ctx.BlockTime())
	require.False(t, iterator.Valid())

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Minute * 10))
	iterator = app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, ctx.BlockTime())
	require.True(t, iterator.Valid())
}

func TestHandler_CreateProposalUpsertDataRegistry_Errors(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name         string
		content      types.Content
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			"Proposer does not have Perm",
			types.NewUpsertDataRegistryProposal(
				"theKey",
				"theHash",
				"theReference",
				"theEncoding",
				1234,
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			errors.Wrap(types.ErrNotEnoughPermissions, types.PermCreateUpsertDataRegistryProposal.String()),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			tt.preparePerms(t, app, ctx)

			handler := gov.NewHandler(app.CustomGovKeeper)
			msg, err := types.NewMsgSubmitProposal(proposerAddr, "title", "some desc", tt.content)
			require.NoError(t, err)
			_, err = handler(ctx, msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_ProposalUpsertDataRegistry(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{
		Time: time.Now(),
	})

	// Set proposer Permissions
	proposerActor := types.NewDefaultActor(proposerAddr)
	err2 := app.CustomGovKeeper.AddWhitelistPermission(ctx, proposerActor, types.PermCreateUpsertDataRegistryProposal)
	require.NoError(t, err2)

	properties := app.CustomGovKeeper.GetNetworkProperties(ctx)
	properties.MinimumProposalEndTime = 10
	app.CustomGovKeeper.SetNetworkProperties(ctx, properties)

	handler := gov.NewHandler(app.CustomGovKeeper)
	proposal := types.NewUpsertDataRegistryProposal(
		"theKey",
		"theHash",
		"theReference",
		"theEncoding",
		1234,
	)
	msg, err := types.NewMsgSubmitProposal(proposerAddr, "title", "some desc", proposal)
	require.NoError(t, err)
	res, err := handler(
		ctx, msg,
	)
	require.NoError(t, err)

	expData, _ := proto.Marshal(&types.MsgSubmitProposalResponse{ProposalID: 1})
	require.Equal(t, expData, res.Data)

	savedProposal, found := app.CustomGovKeeper.GetProposal(ctx, 1)
	require.True(t, found)

	expectedSavedProposal, err := types.NewProposal(
		1,
		"title",
		"some desc",
		types.NewUpsertDataRegistryProposal(
			"theKey",
			"theHash",
			"theReference",
			"theEncoding",
			1234,
		),
		ctx.BlockTime(),
		ctx.BlockTime().Add(time.Second*time.Duration(properties.MinimumProposalEndTime)),
		ctx.BlockTime().Add(time.Second*time.Duration(properties.ProposalEnactmentTime)+time.Second*time.Duration(properties.MinimumProposalEndTime)),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)
	require.Equal(t, expectedSavedProposal, savedProposal)

	// Next proposal ID is increased.
	id := app.CustomGovKeeper.GetNextProposalID(ctx)
	require.Equal(t, uint64(2), id)

	// Is not on finished active proposals.
	iterator := app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, ctx.BlockTime())
	require.False(t, iterator.Valid())

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Minute * 10))
	iterator = app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, ctx.BlockTime())
	require.True(t, iterator.Valid())
}

func TestHandler_VoteProposal_Errors(t *testing.T) {
	voterAddr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	tests := []struct {
		name         string
		msg          *types.MsgVoteProposal
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			"voting time has finished",
			types.NewMsgVoteProposal(
				1, voterAddr, types.OptionAbstain, sdk.ZeroDec(),
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				actor := types.NewNetworkActor(
					voterAddr,
					[]uint64{},
					types.Active,
					[]types.VoteOption{},
					types.NewPermissions(nil, nil),
					1,
				)
				app.CustomGovKeeper.SaveNetworkActor(ctx, actor)

				err = app.CustomGovKeeper.AddWhitelistPermission(ctx, actor, types.PermVoteWhitelistAccountPermissionProposal)
				require.NoError(t, err)

				// Create proposal
				proposal, err := types.NewProposal(
					1,
					"title",
					"some desc",
					types.NewWhitelistAccountPermissionProposal(
						voterAddr,
						types.PermClaimCouncilor,
					),
					ctx.BlockTime(),
					ctx.BlockTime().Add(time.Second*9),
					ctx.BlockTime().Add(time.Second*20),
					ctx.BlockHeight()+2,
					ctx.BlockHeight()+3,
				)

				require.NoError(t, err)
				app.CustomGovKeeper.SaveProposal(ctx, proposal)
			},
			types.ErrVotingTimeEnded,
		},
		{
			"Voter does not have permission to vote this proposal: Assign Permission",
			types.NewMsgVoteProposal(
				1, voterAddr, types.OptionAbstain, sdk.ZeroDec(),
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				actor := types.NewNetworkActor(
					voterAddr,
					[]uint64{},
					types.Active,
					[]types.VoteOption{},
					types.NewPermissions(nil, nil),
					1,
				)
				app.CustomGovKeeper.SaveNetworkActor(ctx, actor)

				// Create proposal
				proposal, err := types.NewProposal(
					1,
					"title",
					"some desc",
					types.NewWhitelistAccountPermissionProposal(
						voterAddr,
						types.PermClaimCouncilor,
					),
					ctx.BlockTime(),
					ctx.BlockTime().Add(time.Second*20),
					ctx.BlockTime().Add(time.Second*30),
					ctx.BlockHeight()+2,
					ctx.BlockHeight()+3,
				)
				require.NoError(t, err)
				app.CustomGovKeeper.SaveProposal(ctx, proposal)
			},
			fmt.Errorf("%s: not enough permissions", types.PermVoteWhitelistAccountPermissionProposal.String()),
		},
		{
			"Voter does not have permission to vote this proposal: Change Data Registry",
			types.NewMsgVoteProposal(
				1, voterAddr, types.OptionAbstain, sdk.ZeroDec(),
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				actor := types.NewNetworkActor(
					voterAddr,
					[]uint64{},
					types.Active,
					[]types.VoteOption{},
					types.NewPermissions(nil, nil),
					1,
				)
				app.CustomGovKeeper.SaveNetworkActor(ctx, actor)

				// Create proposal
				proposal, err := types.NewProposal(
					1,
					"title",
					"some desc",
					types.NewUpsertDataRegistryProposal(
						"theKey",
						"theHash",
						"theReference",
						"theEncoding",
						1234,
					),
					ctx.BlockTime(),
					ctx.BlockTime().Add(time.Second*20),
					ctx.BlockTime().Add(time.Second*30),
					ctx.BlockHeight()+2,
					ctx.BlockHeight()+3,
				)
				require.NoError(t, err)
				app.CustomGovKeeper.SaveProposal(ctx, proposal)
			},
			fmt.Errorf("%s: not enough permissions", types.PermVoteUpsertDataRegistryProposal.String()),
		},
		{
			"Proposal does not exist",
			types.NewMsgVoteProposal(
				1, voterAddr, types.OptionAbstain, sdk.ZeroDec(),
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				actor := types.NewNetworkActor(
					voterAddr,
					[]uint64{},
					types.Active,
					[]types.VoteOption{},
					types.NewPermissions(nil, nil),
					1,
				)
				app.CustomGovKeeper.SaveNetworkActor(ctx, actor)
				err2 := app.CustomGovKeeper.AddWhitelistPermission(ctx, actor, types.PermVoteWhitelistAccountPermissionProposal)
				require.NoError(t, err2)
			},
			types.ErrProposalDoesNotExist,
		},
		{
			"Voter is not active",
			types.NewMsgVoteProposal(
				1, voterAddr, types.OptionAbstain, sdk.ZeroDec(),
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				actor := types.NewDefaultActor(voterAddr)
				actor.Deactivate()
				app.CustomGovKeeper.SaveNetworkActor(ctx, actor)
				err2 := app.CustomGovKeeper.AddWhitelistPermission(ctx, actor, types.PermVoteWhitelistAccountPermissionProposal)
				require.NoError(t, err2)
			},
			types.ErrActorIsNotActive,
		},
		{
			"Voter does not have permission to vote this proposal: Change Network Property",
			types.NewMsgVoteProposal(
				1, voterAddr, types.OptionAbstain, sdk.ZeroDec(),
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				actor := types.NewNetworkActor(
					voterAddr,
					[]uint64{},
					types.Active,
					[]types.VoteOption{},
					types.NewPermissions(nil, nil),
					1,
				)
				app.CustomGovKeeper.SaveNetworkActor(ctx, actor)

				// Create proposal
				proposal, err := types.NewProposal(
					1,
					"title",
					"some desc",
					types.NewSetNetworkPropertyProposal(
						types.MinTxFee,
						types.NetworkPropertyValue{Value: 1234},
					),
					ctx.BlockTime(),
					ctx.BlockTime().Add(time.Second*20),
					ctx.BlockTime().Add(time.Second*30),
					ctx.BlockHeight()+2,
					ctx.BlockHeight()+3,
				)
				require.NoError(t, err)
				app.CustomGovKeeper.SaveProposal(ctx, proposal)
			},
			fmt.Errorf("%s: not enough permissions", types.PermVoteSetNetworkPropertyProposal.String()),
		},
		{
			"Voter does not have permission to vote this proposal: UpsertTokenAlias",
			types.NewMsgVoteProposal(
				1, voterAddr, types.OptionAbstain, sdk.ZeroDec(),
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				actor := types.NewNetworkActor(
					voterAddr,
					[]uint64{},
					types.Active,
					[]types.VoteOption{},
					types.NewPermissions(nil, nil),
					1,
				)
				app.CustomGovKeeper.SaveNetworkActor(ctx, actor)

				// Create proposal
				proposal, err := types.NewProposal(
					1,
					"title",
					"some desc",
					tokenstypes.NewUpsertTokenAliasProposal(
						"eur",
						"Euro",
						"theIcon",
						12,
						[]string{},
						false,
					),
					ctx.BlockTime(),
					ctx.BlockTime().Add(time.Second*20),
					ctx.BlockTime().Add(time.Second*30),
					ctx.BlockHeight()+2,
					ctx.BlockHeight()+3,
				)
				require.NoError(t, err)
				app.CustomGovKeeper.SaveProposal(ctx, proposal)
			},
			fmt.Errorf("%s: not enough permissions", types.PermVoteUpsertTokenAliasProposal.String()),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{}).WithBlockTime(time.Now())

			tt.preparePerms(t, app, ctx)

			// Add some BlockTime.
			ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Second * 10))

			handler := gov.NewHandler(app.CustomGovKeeper)
			_, err := handler(ctx, tt.msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_VoteProposal(t *testing.T) {
	voterAddr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	// Create Voter as active actor.
	actor := types.NewNetworkActor(
		voterAddr,
		[]uint64{},
		types.Active,
		[]types.VoteOption{},
		types.NewPermissions(nil, nil),
		1,
	)
	app.CustomGovKeeper.SaveNetworkActor(ctx, actor)
	err2 := app.CustomGovKeeper.AddWhitelistPermission(ctx, actor, types.PermVoteWhitelistAccountPermissionProposal)
	require.NoError(t, err2)

	// Create proposal
	proposal, err := types.NewProposal(
		1,
		"title",
		"some desc",
		types.NewWhitelistAccountPermissionProposal(
			voterAddr,
			types.PermClaimCouncilor,
		),
		ctx.BlockTime(),
		ctx.BlockTime().Add(time.Second*1),
		ctx.BlockTime().Add(time.Second*10),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)
	app.CustomGovKeeper.SaveProposal(ctx, proposal)

	msg := types.NewMsgVoteProposal(proposal.ProposalId, voterAddr, types.OptionAbstain, sdk.ZeroDec())
	handler := gov.NewHandler(app.CustomGovKeeper)
	_, err = handler(ctx, msg)
	require.NoError(t, err)

	vote, found := app.CustomGovKeeper.GetVote(ctx, proposal.ProposalId, voterAddr)
	require.True(t, found)
	require.Equal(t, types.NewVote(proposal.ProposalId, voterAddr, types.OptionAbstain, sdk.ZeroDec()), vote)
}

func setPermissionToAddr(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context, addr sdk.AccAddress, perm types.PermValue) error {
	proposerActor := types.NewDefaultActor(addr)
	err := app.CustomGovKeeper.AddWhitelistPermission(ctx, proposerActor, perm)
	require.NoError(t, err)

	return nil
}

func TestHandler_CreateProposalSetNetworkProperty_Errors(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name         string
		content      types.Content
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			"Proposer does not have Perm",
			types.NewSetNetworkPropertyProposal(
				types.MaxTxFee,
				types.NetworkPropertyValue{Value: 100000},
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			errors.Wrap(types.ErrNotEnoughPermissions, types.PermCreateSetNetworkPropertyProposal.String()),
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			tt.preparePerms(t, app, ctx)

			handler := gov.NewHandler(app.CustomGovKeeper)
			msg, err := types.NewMsgSubmitProposal(proposerAddr, "title", "some desc", tt.content)
			require.NoError(t, err)
			_, err = handler(ctx, msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_ProposalSetNetworkProperty(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{
		Time: time.Now(),
	})

	// Set proposer Permissions
	proposerActor := types.NewDefaultActor(proposerAddr)
	err2 := app.CustomGovKeeper.AddWhitelistPermission(ctx, proposerActor, types.PermCreateSetNetworkPropertyProposal)
	require.NoError(t, err2)

	properties := app.CustomGovKeeper.GetNetworkProperties(ctx)
	properties.MinimumProposalEndTime = 10
	app.CustomGovKeeper.SetNetworkProperties(ctx, properties)

	handler := gov.NewHandler(app.CustomGovKeeper)
	proposal := types.NewSetNetworkPropertyProposal(
		types.MinTxFee,
		types.NetworkPropertyValue{Value: 1234},
	)
	msg, err := types.NewMsgSubmitProposal(proposerAddr, "title", "some desc", proposal)
	require.NoError(t, err)
	res, err := handler(
		ctx,
		msg,
	)
	require.NoError(t, err)

	expData, _ := proto.Marshal(&types.MsgSubmitProposalResponse{ProposalID: 1})
	require.Equal(t, expData, res.Data)

	savedProposal, found := app.CustomGovKeeper.GetProposal(ctx, 1)
	require.True(t, found)

	expectedSavedProposal, err := types.NewProposal(
		1,
		"title",
		"some desc",
		types.NewSetNetworkPropertyProposal(
			types.MinTxFee,
			types.NetworkPropertyValue{Value: 1234},
		),
		ctx.BlockTime(),
		ctx.BlockTime().Add(time.Second*time.Duration(properties.MinimumProposalEndTime)),
		ctx.BlockTime().Add(time.Second*time.Duration(properties.ProposalEnactmentTime)+time.Second*time.Duration(properties.MinimumProposalEndTime)),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)
	require.Equal(t, expectedSavedProposal, savedProposal)

	// Next proposal ID is increased.
	id := app.CustomGovKeeper.GetNextProposalID(ctx)
	require.Equal(t, uint64(2), id)

	// Is not on finished active proposals.
	iterator := app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, ctx.BlockTime())
	require.False(t, iterator.Valid())

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Minute * 10))
	iterator = app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, ctx.BlockTime())
	require.True(t, iterator.Valid())
}

func TestHandler_CreateProposalCreateRole_Errors(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	tests := []struct {
		name         string
		content      types.Content
		preparePerms func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context)
		expectedErr  error
	}{
		{
			"Proposer does not have Perm",
			types.NewCreateRoleProposal(
				"role1",
				"role1 description",
				[]types.PermValue{},
				[]types.PermValue{types.PermClaimValidator},
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {},
			errors.Wrap(types.ErrNotEnoughPermissions, types.PermCreateRoleProposal.String()),
		},
		{
			"role already exist",
			types.NewCreateRoleProposal(
				"role1",
				"role1 description",
				[]types.PermValue{types.PermClaimCouncilor},
				[]types.PermValue{},
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				proposerActor := types.NewDefaultActor(proposerAddr)
				err := app.CustomGovKeeper.AddWhitelistPermission(
					ctx,
					proposerActor,
					types.PermCreateRoleProposal,
				)
				require.NoError(t, err)

				app.CustomGovKeeper.SetRole(ctx, types.Role{
					Id:          1,
					Sid:         "role1",
					Description: "role1 description",
				})
			},
			types.ErrRoleExist,
		},
		{
			"permissions are empty",
			types.NewCreateRoleProposal(
				"role1000",
				"role1000 description",
				[]types.PermValue{},
				[]types.PermValue{},
			),
			func(t *testing.T, app *simapp.TsukiApp, ctx sdk.Context) {
				proposerActor := types.NewDefaultActor(proposerAddr)
				err := app.CustomGovKeeper.AddWhitelistPermission(
					ctx,
					proposerActor,
					types.PermCreateRoleProposal,
				)
				require.NoError(t, err)
			},
			types.ErrEmptyPermissions,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			app := simapp.Setup(false)
			ctx := app.NewContext(false, tmproto.Header{})

			tt.preparePerms(t, app, ctx)

			handler := gov.NewHandler(app.CustomGovKeeper)
			msg, err := types.NewMsgSubmitProposal(proposerAddr, "title", "some desc", tt.content)
			require.NoError(t, err)
			_, err = handler(ctx, msg)
			require.EqualError(t, err, tt.expectedErr.Error())
		})
	}
}

func TestHandler_ProposalCreateRole(t *testing.T) {
	proposerAddr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	require.NoError(t, err)

	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{
		Time: time.Now(),
	})

	// Set proposer Permissions
	proposerActor := types.NewDefaultActor(proposerAddr)
	err2 := app.CustomGovKeeper.AddWhitelistPermission(ctx, proposerActor, types.PermCreateRoleProposal)
	require.NoError(t, err2)

	properties := app.CustomGovKeeper.GetNetworkProperties(ctx)
	properties.MinimumProposalEndTime = 10
	app.CustomGovKeeper.SetNetworkProperties(ctx, properties)

	handler := gov.NewHandler(app.CustomGovKeeper)
	proposal := types.NewCreateRoleProposal(
		"role1000",
		"role1000 description",
		[]types.PermValue{
			types.PermClaimValidator,
		},
		[]types.PermValue{
			types.PermChangeTxFee,
		},
	)
	msg, err := types.NewMsgSubmitProposal(proposerAddr, "title", "some desc", proposal)
	require.NoError(t, err)
	res, err := handler(
		ctx,
		msg,
	)
	require.NoError(t, err)

	expData, _ := proto.Marshal(&types.MsgSubmitProposalResponse{ProposalID: 1})
	require.Equal(t, expData, res.Data)

	savedProposal, found := app.CustomGovKeeper.GetProposal(ctx, 1)
	require.True(t, found)

	expectedSavedProposal, err := types.NewProposal(
		1,
		"title",
		"some desc",
		types.NewCreateRoleProposal(
			"role1000",
			"role1000 description",
			[]types.PermValue{
				types.PermClaimValidator,
			},
			[]types.PermValue{
				types.PermChangeTxFee,
			},
		),
		ctx.BlockTime(),
		ctx.BlockTime().Add(time.Second*time.Duration(properties.MinimumProposalEndTime)),
		ctx.BlockTime().Add(time.Second*time.Duration(properties.ProposalEnactmentTime)+time.Second*time.Duration(properties.MinimumProposalEndTime)),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)
	require.Equal(t, expectedSavedProposal, savedProposal)

	// Next proposal ID is increased.
	id := app.CustomGovKeeper.GetNextProposalID(ctx)
	require.Equal(t, uint64(2), id)

	// Is not on finished active proposals.
	iterator := app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, ctx.BlockTime())
	require.False(t, iterator.Valid())

	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(time.Minute * 10))
	iterator = app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, ctx.BlockTime())
	require.True(t, iterator.Valid())
}

func TestHandler_SetProposalDurationsProposal(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{
		Time: time.Now(),
	})

	proposal := types.NewSetProposalDurationsProposal(
		[]string{tsukitypes.ProposalTypeCreateRole, tsukitypes.ProposalTypeSetNetworkProperty},
		[]uint64{1200, 2400}, // 20 min, 40min
	)

	router := app.CustomGovKeeper.GetProposalRouter()
	err := router.ApplyProposal(ctx, 1, proposal, sdk.ZeroDec())
	require.NoError(t, err)

	duration := app.CustomGovKeeper.GetProposalDuration(ctx, tsukitypes.ProposalTypeCreateRole)
	require.Equal(t, duration, uint64(1200))

	duration = app.CustomGovKeeper.GetProposalDuration(ctx, tsukitypes.ProposalTypeSetNetworkProperty)
	require.Equal(t, duration, uint64(2400))
}
