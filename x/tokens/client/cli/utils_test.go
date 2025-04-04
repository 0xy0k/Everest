package cli_test

import (
	"fmt"

	"github.com/TsukiCore/tsuki/x/gov/client/cli"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	customstakingcli "github.com/TsukiCore/tsuki/x/staking/client/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s IntegrationTestSuite) WhitelistPermissions(addr sdk.AccAddress, perm govtypes.PermValue) {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetTxSetWhitelistPermissions()
	_, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=%s", customstakingcli.FlagAddr, addr.String()),
		fmt.Sprintf("--%s=%d", cli.FlagPermission, perm),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
	})
	s.Require().NoError(err)

	err = s.network.WaitForNextBlock()
	s.Require().NoError(err)

	// We check if the user has the permissions
	cmd = cli.GetCmdQueryPermissions()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		addr.String(),
	})
	s.Require().NoError(err)

}
