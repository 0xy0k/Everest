package cli_test

import (
	"context"
	"fmt"
	"strings"

	stakingcli "github.com/TsukiCore/tsuki/x/staking/client/cli"

	"github.com/TsukiCore/tsuki/x/gov/client/cli"
	customgovtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s IntegrationTestSuite) TestWhitelistRolePermission() {
	// Query permissions for role Validator
	val := s.network.Validators[0]

	cmd := cli.GetCmdQueryRolePermissions()
	_, out := testutil.ApplyMockIO(cmd)
	clientCtx := val.ClientCtx.WithOutput(out).WithOutputFormat("json")

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	cmd.SetArgs([]string{
		"0", // RoleInTest
	})

	err := cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	var perms customgovtypes.Permissions
	val.ClientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &perms)

	s.Require().False(perms.IsWhitelisted(customgovtypes.PermSetPermissions))

	out.Reset()

	// Send Tx To Whitelist permission
	cmd = cli.GetTxWhitelistRolePermission()
	cmd.SetArgs([]string{
		"0", // Role created in test
		"1", // PermSetPermission
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(100))).String()),
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	// Query again to check if it has the new permission
	out.Reset()

	cmd = cli.GetCmdQueryRolePermissions()

	cmd.SetArgs([]string{
		"0", // RoleCreatedInTest
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	var newPerms customgovtypes.Permissions
	val.ClientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &newPerms)

	s.Require().True(newPerms.IsWhitelisted(customgovtypes.PermSetPermissions))
}

func (s IntegrationTestSuite) TestBlacklistRolePermission() {
	// Query permissions for role Validator
	val := s.network.Validators[0]

	cmd := cli.GetCmdQueryRolePermissions()
	_, out := testutil.ApplyMockIO(cmd)
	clientCtx := val.ClientCtx.WithOutput(out).WithOutputFormat("json")

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	cmd.SetArgs([]string{
		"2", // RoleValidator
	})

	err := cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	var perms customgovtypes.Permissions
	val.ClientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &perms)

	s.Require().True(perms.IsWhitelisted(customgovtypes.PermClaimValidator))
	s.Require().False(perms.IsBlacklisted(customgovtypes.PermClaimCouncilor))

	// Send Tx To Blacklist permission
	out.Reset()

	cmd = cli.GetTxBlacklistRolePermission()
	cmd.SetArgs([]string{
		"2", // RoleValidator
		"3", // PermClaimCouncilor
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(100))).String()),
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	// Query again to check if it has the new permission
	out.Reset()

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	cmd = cli.GetCmdQueryRolePermissions()

	cmd.SetArgs([]string{
		"2", // RoleValidator
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	var newPerms customgovtypes.Permissions
	val.ClientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &newPerms)

	s.Require().True(newPerms.IsWhitelisted(customgovtypes.PermClaimValidator))
	s.Require().True(newPerms.IsBlacklisted(customgovtypes.PermClaimCouncilor))
}

func (s IntegrationTestSuite) TestRemoveWhitelistRolePermission() {
	// Query permissions for role Validator
	val := s.network.Validators[0]

	cmd := cli.GetCmdQueryRolePermissions()
	_, out := testutil.ApplyMockIO(cmd)
	clientCtx := val.ClientCtx.WithOutput(out).WithOutputFormat("json")

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	cmd.SetArgs([]string{
		"0", // RoleInTest
	})

	err := cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	var perms customgovtypes.Permissions
	val.ClientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &perms)

	s.Require().True(perms.IsWhitelisted(customgovtypes.PermClaimValidator))

	// Send Tx To Blacklist permission
	out.Reset()

	cmd = cli.GetTxRemoveWhitelistRolePermission()
	cmd.SetArgs([]string{
		"0", // RoleValidator
		"2", // PermClaimValidator
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(100))).String()),
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	// Query again to check if it has the new permission
	out.Reset()

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	cmd = cli.GetCmdQueryRolePermissions()

	cmd.SetArgs([]string{
		"0", // RoleInTest
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	var newPerms customgovtypes.Permissions
	val.ClientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &newPerms)

	s.Require().False(newPerms.IsWhitelisted(customgovtypes.PermClaimValidator))
}

func (s IntegrationTestSuite) TestRemoveBlacklistRolePermission() {
	// Query permissions for role RoleInTest
	val := s.network.Validators[0]

	cmd := cli.GetCmdQueryRolePermissions()
	_, out := testutil.ApplyMockIO(cmd)
	clientCtx := val.ClientCtx.WithOutput(out).WithOutputFormat("json")

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	cmd.SetArgs([]string{
		"0", // RoleInTest
	})

	err := cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	var perms customgovtypes.Permissions
	val.ClientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &perms)

	s.Require().True(perms.IsBlacklisted(customgovtypes.PermClaimCouncilor))

	// Send Tx To Remove Blacklist Permissions
	out.Reset()

	cmd = cli.GetTxRemoveBlacklistRolePermission()
	cmd.SetArgs([]string{
		"0", // RoleValidator
		"3", // PermClaimCouncilor
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(100))).String()),
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	// Query again to check if it has the new permission
	out.Reset()

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	cmd = cli.GetCmdQueryRolePermissions()

	cmd.SetArgs([]string{
		"0", // RoleInTest
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	var newPerms customgovtypes.Permissions
	val.ClientCtx.JSONMarshaler.MustUnmarshalJSON(out.Bytes(), &newPerms)

	s.Require().False(newPerms.IsBlacklisted(customgovtypes.PermClaimCouncilor))
}

func (s IntegrationTestSuite) TestCreateRole() {
	// Query permissions for role Non existing role yet
	val := s.network.Validators[0]

	cmd := cli.GetCmdQueryRolePermissions()
	_, out := testutil.ApplyMockIO(cmd)
	clientCtx := val.ClientCtx.WithOutput(out).WithOutputFormat("json")

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	cmd.SetArgs([]string{
		"1234", // RoleInTest
	})

	err := cmd.ExecuteContext(ctx)
	s.Require().Error(err)
	strings.Contains(err.Error(), customgovtypes.ErrRoleDoesNotExist.Error())

	// Add role
	out.Reset()

	cmd = cli.GetTxCreateRole()
	cmd.SetArgs([]string{
		"1234", // RoleValidator
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(100))).String()),
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	// Query again the role
	out.Reset()
	cmd = cli.GetCmdQueryRolePermissions()

	cmd.SetArgs([]string{
		"1234", // RoleInTest
	})

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)
}

func (s IntegrationTestSuite) TestAssignRoles_AndRemoveRoles() {
	val := s.network.Validators[0]

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	s.Require().NoError(err)

	cmd := cli.GetTxAssignRole()
	_, out := testutil.ApplyMockIO(cmd)
	clientCtx := val.ClientCtx.WithOutput(out).WithOutputFormat("json")

	ctx := context.Background()
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)

	cmd.SetArgs(
		[]string{
			"0", // Role created in test
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			fmt.Sprintf("--%s=%s", stakingcli.FlagAddr, addr),
			fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(100))).String()),
		},
	)

	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	roles := GetRolesByAddress(s.T(), s.network, addr)
	s.Require().Equal([]uint64{uint64(customgovtypes.RoleUndefined)}, roles)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	out.Reset()

	cmd = cli.GetTxRemoveRole()
	cmd.SetArgs(
		[]string{
			"0", // Role created in test
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			fmt.Sprintf("--%s=%s", stakingcli.FlagAddr, addr),
			fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(100))).String()),
		},
	)
	err = cmd.ExecuteContext(ctx)
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	roles = GetRolesByAddress(s.T(), s.network, addr)
	s.Require().Equal([]uint64{}, roles)
}

func (s IntegrationTestSuite) TestGetRolesByAddress() {
	val := s.network.Validators[0]

	roles := GetRolesByAddress(s.T(), s.network, val.Address)

	s.Require().Equal([]uint64{uint64(customgovtypes.RoleSudo)}, roles)
}
