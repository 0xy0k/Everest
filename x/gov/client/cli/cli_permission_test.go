package cli_test

import (
	"fmt"

	"github.com/TsukiCore/tsuki/x/gov/client/cli"
	"github.com/TsukiCore/tsuki/x/gov/types"
	stakingcli "github.com/TsukiCore/tsuki/x/staking/client/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s IntegrationTestSuite) TestGetTxSetWhitelistPermissions() {
	val := s.network.Validators[0]
	cmd := cli.GetTxSetWhitelistPermissions()

	// We create some random address where we will give perms.
	addr, err := sdk.AccAddressFromBech32("tsuki15ky9du8a2wlstz6fpx3p4mqpjyrm5cgqzp4f3d")
	s.Require().NoError(err)

	clientCtx := val.ClientCtx.WithOutputFormat("json")
	out, err := clitestutil.ExecTestCLICmd(
		clientCtx,
		cmd,
		[]string{
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			fmt.Sprintf("--%s=%s", stakingcli.FlagAddr, addr.String()),
			fmt.Sprintf("--%s=%s", cli.FlagPermission, "1"),
			fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
		},
	)
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	// We check if the user has the permissions
	cmd = cli.GetCmdQueryPermissions()

	out, err = clitestutil.ExecTestCLICmd(
		clientCtx,
		cmd,
		[]string{
			addr.String(),
		},
	)
	s.Require().NoError(err)

	var perms types.Permissions
	clientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &perms)
}

func (s IntegrationTestSuite) TestGetTxSetBlacklistPermissions() {
	val := s.network.Validators[0]

	// We create some random address where we will give perms.
	addr, err := sdk.AccAddressFromBech32("tsuki1alzyfq40zjsveat87jlg8jxetwqmr0a29sgd0f")
	s.Require().NoError(err)

	clientCtx := val.ClientCtx.WithOutputFormat("json")

	out, err := clitestutil.ExecTestCLICmd(
		clientCtx,
		cli.GetTxSetBlacklistPermissions(),
		[]string{
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			fmt.Sprintf("--%s=%s", stakingcli.FlagAddr, addr.String()),
			fmt.Sprintf("--%s=%s", cli.FlagPermission, "1"),
			fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
		},
	)
	s.Require().NoError(err)
	s.T().Logf("error %s", out.String())

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

}

func (s IntegrationTestSuite) TestGetTxSetWhitelistPermissions_WithUserThatDoesNotHaveSetPermissions() {
	val := s.network.Validators[0]

	// We create some random address where we will give perms.
	newAccount, _, err := val.ClientCtx.Keyring.NewMnemonic("test", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	s.Require().NoError(err)

	addr, err := newAccount.GetAddress()
	s.Require().NoError(err)
	s.SendValue(val.ClientCtx, val.Address, addr, sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100)))

	// Now we try to set permissions with a user that does not have.
	cmd := cli.GetTxSetWhitelistPermissions()
	clientCtx := val.ClientCtx
	_, _ = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, addr.String()),
		fmt.Sprintf("--%s=%s", stakingcli.FlagAddr, val.Address.String()),
		fmt.Sprintf("--%s=%s", cli.FlagPermission, "1"),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
	})
}
