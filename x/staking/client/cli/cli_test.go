package cli_test

import (
	"fmt"
	"testing"

	customgovcli "github.com/TsukiCore/tsuki/x/gov/client/cli"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/cosmos/cosmos-sdk/client/flags"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	dbm "github.com/cometbft/cometbft-db"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/TsukiCore/tsuki/app"
	simapp "github.com/TsukiCore/tsuki/app"
	appparams "github.com/TsukiCore/tsuki/app/params"
	"github.com/TsukiCore/tsuki/testutil/network"
	"github.com/TsukiCore/tsuki/x/staking/client/cli"
	customtypes "github.com/TsukiCore/tsuki/x/staking/types"
	pruningtypes "github.com/cosmos/cosmos-sdk/store/pruning/types"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	appparams.SetConfig()
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	encodingConfig := simapp.MakeEncodingConfig()
	cfg.Codec = encodingConfig.Marshaler
	cfg.TxConfig = encodingConfig.TxConfig

	cfg.NumValidators = 1

	cfg.AppConstructor = func(val network.Validator, chainId string) servertypes.Application {
		return app.NewInitApp(
			val.Ctx.Logger, dbm.NewMemDB(), nil, true, make(map[int64]bool), val.Ctx.Config.RootDir, 0,
			simapp.MakeEncodingConfig(),
			simapp.EmptyAppOptions{},
			baseapp.SetPruning(pruningtypes.NewPruningOptionsFromString(val.AppConfig.Pruning)),
			baseapp.SetMinGasPrices(val.AppConfig.MinGasPrices),
			baseapp.SetChainID(chainId),
		)
	}

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err := s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestQueryValidator() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	cmd := cli.GetCmdQueryValidator()
	out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", cli.FlagValAddr, val.ValAddress.String()),
	})
	s.Require().NoError(err)

	var respValidator customtypes.Validator
	clientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &respValidator)

	s.Require().Equal(val.ValAddress, respValidator.ValKey)

	var pubkey cryptotypes.PubKey
	err = s.cfg.Codec.UnpackAny(respValidator.PubKey, &pubkey)
	s.Require().NoError(err)
	s.Require().Equal(val.PubKey, pubkey)

	// Query by Acc Addrs.
	cmd = cli.GetCmdQueryValidator()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", cli.FlagAddr, val.Address.String()),
	})
	s.Require().NoError(err)

	clientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &respValidator)

	s.Require().Equal(val.ValAddress, respValidator.ValKey)

	err = s.cfg.Codec.UnpackAny(respValidator.PubKey, &pubkey)
	s.Require().NoError(err)
	s.Require().Equal(val.PubKey, pubkey)

	// Query by moniker.
	cmd = cli.GetCmdQueryValidator()
	out, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", cli.FlagMoniker, val.Moniker),
	})
	s.Require().NoError(err)

	clientCtx.Codec.MustUnmarshalJSON(out.Bytes(), &respValidator)

	s.Require().Equal(val.ValAddress, respValidator.ValKey)

	err = s.cfg.Codec.UnpackAny(respValidator.PubKey, &pubkey)
	s.Require().NoError(err)
	s.Require().Equal(val.PubKey, pubkey)
}

func (s *IntegrationTestSuite) TestQueryValidator_Errors() {
	val := s.network.Validators[0]
	clientCtx := val.ClientCtx

	nonExistingAddr, err := sdk.ValAddressFromBech32("tsukivaloper15ky9du8a2wlstz6fpx3p4mqpjyrm5cgpv3al5n")
	s.Require().NoError(err)

	cmd := cli.GetCmdQueryValidator()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", cli.FlagValAddr, nonExistingAddr.String()),
	})
	s.Require().EqualError(err, "rpc error: code = Unknown desc = validator not found: key not found: unknown request")

	// Non existing moniker.
	cmd = cli.GetCmdQueryValidator()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", cli.FlagAddr, sdk.AccAddress(nonExistingAddr).String()),
	})
	s.Require().EqualError(err, "rpc error: code = Unknown desc = validator not found: key not found: unknown request")

	// Non existing moniker.
	cmd = cli.GetCmdQueryValidator()
	_, err = clitestutil.ExecTestCLICmd(clientCtx, cmd, []string{
		fmt.Sprintf("--%s=%s", cli.FlagMoniker, "weirdMoniker"),
	})
	s.Require().EqualError(err, "rpc error: code = Unknown desc = validator with moniker weirdMoniker not found: key not found: unknown request")
}

func (s IntegrationTestSuite) TestCreateProposalUnjailValidator() {
	// Query permissions for role Validator
	val := s.network.Validators[0]

	clientCtx := val.ClientCtx.WithOutputFormat("json")
	out, err := clitestutil.ExecTestCLICmd(
		clientCtx,
		cli.GetTxProposalUnjailValidatorCmd(),
		[]string{
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			fmt.Sprintf("--%s=%s", cli.FlagTitle, "title"),
			fmt.Sprintf("--%s=%s", cli.FlagDescription, "some desc"),
			fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
			val.ValAddress.String(),
			"theReference",
		},
	)
	s.Require().NoError(err)
	fmt.Printf("%s", out.String())

	// Vote Proposal
	out, err = clitestutil.ExecTestCLICmd(
		clientCtx,
		customgovcli.GetTxVoteProposal(),
		[]string{
			fmt.Sprintf("%d", 1), // Proposal ID
			fmt.Sprintf("%d", govtypes.OptionYes),
			fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
			fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
			fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
			fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.DefaultDenom, sdk.NewInt(100))).String()),
		},
	)
	s.Require().NoError(err)
	fmt.Printf("%s", out.String())
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
