package cli_test

import (
	"fmt"
	"strings"

	"github.com/TsukiCore/tsuki/x/gov/client/cli"
	"github.com/TsukiCore/tsuki/x/gov/types"
	stakingcli "github.com/TsukiCore/tsuki/x/staking/client/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s IntegrationTestSuite) TestWhitelistRolePermission() {
	// Query permissions for role Validator
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetCmdQueryRole()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // RoleInTest
	})
	s.Require().NoError(err)

	var roleQuery types.RoleQuery
	val.ClientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &roleQuery)
	s.Require().False(roleQuery.Permissions.IsWhitelisted(types.PermSetPermissions))

	// Send Tx To Whitelist permission
	cmd = cli.GetTxWhitelistRolePermission()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // Role created in test
		"1", // PermSetPermission
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
	})
	s.Require().NoError(err)

	// Query again to check if it has the new permission
	cmd = cli.GetCmdQueryRole()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // RoleCreatedInTest
	})
	s.Require().NoError(err)

	var newRoleQuery types.RoleQuery
	val.ClientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &newRoleQuery)
}

func (s IntegrationTestSuite) TestBlacklistRolePermission() {
	// Query permissions for role Validator
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetCmdQueryRole()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // RoleValidator
	})
	s.Require().NoError(err)

	var roleQuery types.RoleQuery
	val.ClientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &roleQuery)
	s.Require().True(roleQuery.Permissions.IsWhitelisted(types.PermClaimValidator))
	s.Require().False(roleQuery.Permissions.IsBlacklisted(types.PermClaimCouncilor))

	// Send Tx To Blacklist permission
	cmd = cli.GetTxBlacklistRolePermission()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // RoleValidator
		"3", // PermClaimCouncilor
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
	})
	s.Require().NoError(err)

	// Query again to check if it has the new permission
	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	cmd = cli.GetCmdQueryRole()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // RoleValidator
	})
	s.Require().NoError(err)

	var newRoleQuery types.RoleQuery
	val.ClientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &newRoleQuery)
	s.Require().True(newRoleQuery.Permissions.IsWhitelisted(types.PermClaimValidator))
}

func (s IntegrationTestSuite) TestRemoveWhitelistRolePermission() {
	// Query permissions for role Validator
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetCmdQueryRole()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // RoleInTest
	})
	s.Require().NoError(err)

	var roleQuery types.RoleQuery
	val.ClientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &roleQuery)
	s.Require().True(roleQuery.Permissions.IsWhitelisted(types.PermClaimValidator))

	// Send Tx To Blacklist permission
	cmd = cli.GetTxRemoveWhitelistRolePermission()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // RoleValidator
		"2", // PermClaimValidator
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
	})
	s.Require().NoError(err)

	// Query again to check if it has the new permission
	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	cmd = cli.GetCmdQueryRole()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // RoleInTest
	})
	s.Require().NoError(err)

	var newRoleQuery types.RoleQuery
	val.ClientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &newRoleQuery)
	s.Require().False(newRoleQuery.Permissions.IsWhitelisted(types.PermClaimValidator))
}

func (s IntegrationTestSuite) TestRemoveBlacklistRolePermission() {
	// Query permissions for role RoleInTest
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetCmdQueryRole()
	_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"sudo",
	})
	s.Require().NoError(err)

	// Send Tx To Remove Blacklist Permissions
	cmd = cli.GetTxRemoveBlacklistRolePermission()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"sudo", // RoleValidator
		"3",    // PermClaimCouncilor
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
	})
	s.Require().NoError(err)

	// Query again to check if it has the new permission
	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)
}

func (s IntegrationTestSuite) TestCreateRole() {
	// Query permissions for role Non existing role yet
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetCmdQueryRole()
	_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"myRole", // RoleInTest
	})
	s.Require().Error(err)
	strings.Contains(err.Error(), types.ErrRoleDoesNotExist.Error())

	// Add role
	cmd = cli.GetTxCreateRole()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"myRole", "myRole", // RoleValidator
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
	})
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

}

func (s IntegrationTestSuite) TestAssignRoles_AndUnassignRoles() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	s.Require().NoError(err)

	cmd := cli.GetTxAssignRole()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		"2", // Role created in test
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=%s", stakingcli.FlagAddr, addr),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
	})
	s.Require().NoError(err)

}

func (s IntegrationTestSuite) TestGetRolesByAddress() {
	val := s.network.Validators[0]

	roles := GetRolesByAddress(s.T(), s.network, val.Address)

	s.Require().Equal([]uint64{uint64(types.RoleSudo)}, roles)
}
