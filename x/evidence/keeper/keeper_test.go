package keeper_test

import (
	"encoding/hex"
	"fmt"
	"time"

	simapp "github.com/TsukiCore/tsuki/app"
	appparams "github.com/TsukiCore/tsuki/app/params"
	"github.com/TsukiCore/tsuki/x/evidence/exported"
	"github.com/TsukiCore/tsuki/x/evidence/keeper"
	"github.com/TsukiCore/tsuki/x/evidence/types"
	"github.com/TsukiCore/tsuki/x/staking"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/suite"
)

var (
	pubkeys = []cryptotypes.PubKey{
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB50"),
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB51"),
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB52"),
	}

	valAddresses = []sdk.ValAddress{
		sdk.ValAddress(pubkeys[0].Address()),
		sdk.ValAddress(pubkeys[1].Address()),
		sdk.ValAddress(pubkeys[2].Address()),
	}

	initAmt   = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
	initCoins = sdk.NewCoins(sdk.NewCoin(appparams.DefaultDenom, initAmt))
)

func newPubKey(pk string) (res cryptotypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}

	pubkey := &ed25519.PubKey{Key: pkBytes}

	return pubkey
}

func testEquivocationHandler(_ interface{}) types.Handler {
	return func(ctx sdk.Context, e exported.Evidence) error {
		if err := e.ValidateBasic(); err != nil {
			return err
		}

		ee, ok := e.(*types.Equivocation)
		if !ok {
			return fmt.Errorf("unexpected evidence type: %T", e)
		}
		if ee.Height%2 == 0 {
			return fmt.Errorf("unexpected even evidence height: %d", ee.Height)
		}

		return nil
	}
}

type KeeperTestSuite struct {
	suite.Suite

	ctx sdk.Context
	app *simapp.TsukiApp

	queryClient types.QueryClient
	stakingHdl  sdk.Handler
}

func (suite *KeeperTestSuite) SetupTest() {
	checkTx := false
	app := simapp.Setup(checkTx)

	// recreate keeper in order to use custom testing types
	evidenceKeeper := keeper.NewKeeper(
		app.AppCodec(), app.GetKey(types.StoreKey), app.CustomStakingKeeper, app.CustomSlashingKeeper,
	)
	router := types.NewRouter()
	router = router.AddRoute(types.RouteEquivocation, testEquivocationHandler(*evidenceKeeper))
	evidenceKeeper.SetRouter(router)

	app.EvidenceKeeper = *evidenceKeeper

	suite.ctx = app.BaseApp.NewContext(checkTx, tmproto.Header{Height: 1})
	suite.app = app

	for i, addr := range valAddresses {
		addr := sdk.AccAddress(addr)
		app.AccountKeeper.SetAccount(suite.ctx, authtypes.NewBaseAccount(addr, pubkeys[i], uint64(i), 0))
	}

	queryHelper := baseapp.NewQueryServerTestHelper(suite.ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.EvidenceKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
	suite.stakingHdl = staking.NewHandler(app.CustomStakingKeeper, app.CustomGovKeeper)
}

func (suite *KeeperTestSuite) populateEvidence(ctx sdk.Context, numEvidence int) []exported.Evidence {
	evidence := make([]exported.Evidence, numEvidence)

	for i := 0; i < numEvidence; i++ {
		pk := ed25519.GenPrivKey()

		evidence[i] = &types.Equivocation{
			Height:           11,
			Power:            100,
			Time:             time.Now().UTC(),
			ConsensusAddress: sdk.ConsAddress(pk.PubKey().Address().Bytes()).String(),
		}

		suite.Nil(suite.app.EvidenceKeeper.SubmitEvidence(ctx, evidence[i]))
	}

	return evidence
}

func (suite *KeeperTestSuite) populateValidators(ctx sdk.Context) {
	// add accounts
	for _, addr := range valAddresses {
		err := suite.app.BankKeeper.MintCoins(ctx, minttypes.ModuleName, initCoins)
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, sdk.AccAddress(addr), initCoins)
		suite.NoError(err)
	}
}

func (suite *KeeperTestSuite) TestSubmitValidEvidence() {
	ctx := suite.ctx.WithIsCheckTx(false)
	pk := ed25519.GenPrivKey()

	e := &types.Equivocation{
		Height:           1,
		Power:            100,
		Time:             time.Now().UTC(),
		ConsensusAddress: sdk.ConsAddress(pk.PubKey().Address().Bytes()).String(),
	}

	suite.Nil(suite.app.EvidenceKeeper.SubmitEvidence(ctx, e))

	res, ok := suite.app.EvidenceKeeper.GetEvidence(ctx, e.Hash())
	suite.True(ok)
	suite.Equal(e, res)
}

func (suite *KeeperTestSuite) TestSubmitValidEvidence_Duplicate() {
	ctx := suite.ctx.WithIsCheckTx(false)
	pk := ed25519.GenPrivKey()

	e := &types.Equivocation{
		Height:           1,
		Power:            100,
		Time:             time.Now().UTC(),
		ConsensusAddress: sdk.ConsAddress(pk.PubKey().Address().Bytes()).String(),
	}

	suite.Nil(suite.app.EvidenceKeeper.SubmitEvidence(ctx, e))
	suite.Error(suite.app.EvidenceKeeper.SubmitEvidence(ctx, e))

	res, ok := suite.app.EvidenceKeeper.GetEvidence(ctx, e.Hash())
	suite.True(ok)
	suite.Equal(e, res)
}

func (suite *KeeperTestSuite) TestSubmitInvalidEvidence() {
	ctx := suite.ctx.WithIsCheckTx(false)
	pk := ed25519.GenPrivKey()
	e := &types.Equivocation{
		Height:           0,
		Power:            100,
		Time:             time.Now().UTC(),
		ConsensusAddress: sdk.ConsAddress(pk.PubKey().Address().Bytes()).String(),
	}

	suite.Error(suite.app.EvidenceKeeper.SubmitEvidence(ctx, e))

	res, ok := suite.app.EvidenceKeeper.GetEvidence(ctx, e.Hash())
	suite.False(ok)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestIterateEvidence() {
	ctx := suite.ctx.WithIsCheckTx(false)
	numEvidence := 100
	suite.populateEvidence(ctx, numEvidence)

	evidence := suite.app.EvidenceKeeper.GetAllEvidence(ctx)
	suite.Len(evidence, numEvidence)
}

func (suite *KeeperTestSuite) TestGetEvidenceHandler() {
	handler, err := suite.app.EvidenceKeeper.GetEvidenceHandler((&types.Equivocation{}).Route())
	suite.NoError(err)
	suite.NotNil(handler)

	handler, err = suite.app.EvidenceKeeper.GetEvidenceHandler("invalidHandler")
	suite.Error(err)
	suite.Nil(handler)
}
