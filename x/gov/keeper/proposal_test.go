package keeper_test

import (
	"testing"
	"time"

	simapp "github.com/TsukiCore/tsuki/app"
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultProposalIdAtDefaultGenesis(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	proposalID := app.CustomGovKeeper.GetNextProposalID(ctx)
	require.Equal(t, uint64(1), proposalID)
}

func TestKeeper_EncodingContentType(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction))
	addr := addrs[0]

	proposal1, err := types.NewProposal(
		1,
		"title",
		"some desc",
		types.NewWhitelistAccountPermissionProposal(
			addr,
			types.PermSetPermissions,
		),
		time.Now(),
		time.Now().Add(1*time.Second),
		time.Now().Add(10*time.Second),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)

	app.CustomGovKeeper.SaveProposal(ctx, proposal1)

	saveProposal, found := app.CustomGovKeeper.GetProposal(ctx, proposal1.ProposalId)
	require.True(t, found)

	require.Equal(t, proposal1.GetContent(), saveProposal.GetContent())

	content, ok := saveProposal.GetContent().(*types.WhitelistAccountPermissionProposal)
	require.True(t, ok)
	require.Equal(t, addr, content.Address)
	require.Equal(t, types.PermSetPermissions, content.Permission)
}

func TestKeeper_GetProposals(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction))
	addr := addrs[0]

	proposal1, err := types.NewProposal(
		1,
		"title",
		"some desc",
		types.NewWhitelistAccountPermissionProposal(
			addr,
			types.PermSetPermissions,
		),
		time.Now(),
		time.Now().Add(1*time.Second),
		time.Now().Add(10*time.Second),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)

	app.CustomGovKeeper.SaveProposal(ctx, proposal1)

	proposals, err := app.CustomGovKeeper.GetProposals(ctx)
	require.NoError(t, err)
	require.Len(t, proposals, 1)

	proposal2, err := types.NewProposal(
		2,
		"title",
		"some desc",
		types.NewWhitelistAccountPermissionProposal(
			addr,
			types.PermSetPermissions,
		),
		time.Now(),
		time.Now().Add(1*time.Second),
		time.Now().Add(10*time.Second),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	app.CustomGovKeeper.SaveProposal(ctx, proposal2)
	proposals, err = app.CustomGovKeeper.GetProposals(ctx)
	require.NoError(t, err)
	require.Len(t, proposals, 2)
	require.Equal(t, proposals[1].ProposalId, proposal2.ProposalId)
	require.Equal(t, proposals[1].Content, proposal2.Content)
	require.Equal(t, proposals[1].Result, proposal2.Result)
	require.Equal(t, proposals[1].SubmitTime.UTC().String(), proposal2.SubmitTime.UTC().String())
	require.Equal(t, proposals[1].VotingEndTime.UTC().String(), proposal2.VotingEndTime.UTC().String())
	require.Equal(t, proposals[1].EnactmentEndTime.UTC().String(), proposal2.EnactmentEndTime.UTC().String())
}

func TestSaveProposalReturnsTheProposalID_AndIncreasesLast(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	proposalID := app.CustomGovKeeper.GetNextProposalIDAndIncrement(ctx)
	require.Equal(t, uint64(1), proposalID)

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction))
	addr := addrs[0]

	proposal, err := types.NewProposal(
		proposalID,
		"title",
		"some desc",
		types.NewWhitelistAccountPermissionProposal(
			addr,
			types.PermClaimValidator,
		),
		ctx.BlockTime(),
		ctx.BlockTime().Add(10*time.Minute),
		ctx.BlockTime().Add(20*time.Minute),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)
	app.CustomGovKeeper.SaveProposal(ctx, proposal)

	// nextProposalID should be 2
	proposalID = app.CustomGovKeeper.GetNextProposalID(ctx)
	require.Equal(t, uint64(2), proposalID)

	// Get proposal
	savedProposal, found := app.CustomGovKeeper.GetProposal(ctx, proposal.ProposalId)
	require.True(t, found)
	require.Equal(t, proposal, savedProposal)
}

func TestKeeper_SaveVote(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction))
	addr := addrs[0]

	// Vote not saved yet
	_, found := app.CustomGovKeeper.GetVote(ctx, 1, addr)
	require.False(t, found)

	vote := types.NewVote(1, addr, types.OptionAbstain, sdk.ZeroDec())

	app.CustomGovKeeper.SaveVote(ctx, vote)

	savedVote, found := app.CustomGovKeeper.GetVote(ctx, 1, addr)
	require.True(t, found)
	require.Equal(t, vote, savedVote)
}

func TestKeeper_AddProposalToActiveQueue(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction))
	addr := addrs[0]

	baseEndTime := time.Now()
	for _, i := range []uint64{1, 2, 3} {
		endTime := baseEndTime.Add(time.Second * time.Duration(i))

		proposal, err := types.NewProposal(
			i,
			"title",
			"some desc",
			types.NewWhitelistAccountPermissionProposal(
				addr,
				types.PermSetPermissions,
			),
			baseEndTime,
			endTime,
			endTime,
			ctx.BlockHeight()+2,
			ctx.BlockHeight()+3,
		)
		require.NoError(t, err)

		app.CustomGovKeeper.SaveProposal(ctx, proposal)
		app.CustomGovKeeper.AddToActiveProposals(ctx, proposal)
	}

	// We only get until endtime of the second proposal.
	iterator := app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, baseEndTime.Add(2*time.Second))
	defer iterator.Close()
	requireIteratorCount(t, iterator, 2)

	// We remove one ActiveProposal, the first
	proposal, found := app.CustomGovKeeper.GetProposal(ctx, 1)
	require.True(t, found)
	app.CustomGovKeeper.RemoveActiveProposal(ctx, proposal)

	// We then only get 1 proposal.
	iterator = app.CustomGovKeeper.GetActiveProposalsWithFinishedVotingEndTimeIterator(ctx, baseEndTime.Add(2*time.Second))
	defer iterator.Close()
	requireIteratorCount(t, iterator, 1)
}

func TestKeeper_AddProposalToEnactmentQueue(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 1, sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction))
	addr := addrs[0]

	baseEndTime := time.Now()
	for _, i := range []uint64{1, 2, 3} {
		enactmentEndTime := baseEndTime.Add(time.Duration(i) * time.Second)
		proposal, err := types.NewProposal(
			i,
			"title",
			"some desc",
			types.NewWhitelistAccountPermissionProposal(
				addr,
				types.PermSetPermissions,
			),
			baseEndTime,
			baseEndTime,
			enactmentEndTime,
			ctx.BlockHeight()+2,
			ctx.BlockHeight()+3,
		)
		require.NoError(t, err)

		app.CustomGovKeeper.SaveProposal(ctx, proposal)
		app.CustomGovKeeper.AddToEnactmentProposals(ctx, proposal)
	}

	// We only get until endtime of the second proposal.
	iterator := app.CustomGovKeeper.GetEnactmentProposalsWithFinishedEnactmentEndTimeIterator(ctx, baseEndTime.Add(2*time.Second))
	defer iterator.Close()
	requireIteratorCount(t, iterator, 2)

	// We remove one Proposal from the Enactment list, the first
	proposal, found := app.CustomGovKeeper.GetProposal(ctx, 1)
	require.True(t, found)
	app.CustomGovKeeper.RemoveEnactmentProposal(ctx, proposal)

	// We then only get 1 proposal.
	iterator = app.CustomGovKeeper.GetEnactmentProposalsWithFinishedEnactmentEndTimeIterator(ctx, baseEndTime.Add(2*time.Second))
	defer iterator.Close()
	requireIteratorCount(t, iterator, 1)
}

func TestKeeper_GetProposalVotesIterator(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrsIncremental(app, ctx, 2, sdk.TokensFromConsensusPower(10, sdk.DefaultPowerReduction))
	addr1 := addrs[0]
	addr2 := addrs[1]

	proposal1, err := types.NewProposal(
		1,
		"title",
		"some desc",
		types.NewWhitelistAccountPermissionProposal(
			addr1,
			types.PermSetPermissions,
		),
		time.Now(),
		time.Now().Add(1*time.Second),
		time.Now().Add(10*time.Second),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)

	proposal2, err := types.NewProposal(
		2,
		"title",
		"some desc",
		types.NewWhitelistAccountPermissionProposal(
			addr2,
			types.PermClaimCouncilor,
		),
		time.Now(),
		time.Now().Add(1*time.Second),
		time.Now().Add(10*time.Second),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
	)
	require.NoError(t, err)

	app.CustomGovKeeper.SaveProposal(ctx, proposal1)
	app.CustomGovKeeper.SaveProposal(ctx, proposal2)

	// 1st proposal has 2 votes
	vote1 := types.NewVote(proposal1.ProposalId, addr1, types.OptionYes, sdk.ZeroDec())
	vote2 := types.NewVote(proposal1.ProposalId, addr2, types.OptionYes, sdk.ZeroDec())
	app.CustomGovKeeper.SaveVote(ctx, vote1)
	app.CustomGovKeeper.SaveVote(ctx, vote2)

	// 2nd proposal has 1 vote
	v1 := types.NewVote(proposal2.ProposalId, addr1, types.OptionYes, sdk.ZeroDec())
	app.CustomGovKeeper.SaveVote(ctx, v1)

	// We iterate the 1st proposal
	iterator := app.CustomGovKeeper.GetProposalVotesIterator(ctx, proposal1.ProposalId)
	require.True(t, iterator.Valid())
	totalVotes := 0
	for ; iterator.Valid(); iterator.Next() {
		totalVotes++
	}
	require.Equal(t, 2, totalVotes)

	// We iterate the 2nd proposal
	iterator = app.CustomGovKeeper.GetProposalVotesIterator(ctx, proposal2.ProposalId)
	require.True(t, iterator.Valid())
	totalVotes = 0
	for ; iterator.Valid(); iterator.Next() {
		totalVotes++
	}
	require.Equal(t, 1, totalVotes)
}

func TestKeeper_ProposalDuration(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.NewContext(false, tmproto.Header{})
	ctx = ctx.WithBlockTime(time.Now())

	properties := app.CustomGovKeeper.GetNetworkProperties(ctx)

	// test WhitelistAccountPermissionProposal
	proposalID, err := app.CustomGovKeeper.CreateAndSaveProposalWithContent(ctx, "title", "description", &types.WhitelistAccountPermissionProposal{})
	require.NoError(t, err)

	proposal, found := app.CustomGovKeeper.GetProposal(ctx, proposalID)
	require.True(t, found)

	require.Equal(t, proposal.VotingEndTime.Unix(), ctx.BlockTime().Unix()+int64(properties.MinimumProposalEndTime))

	// test SetNetworkPropertyProposal
	proposalID, err = app.CustomGovKeeper.CreateAndSaveProposalWithContent(ctx, "title", "description", &types.SetNetworkPropertyProposal{})
	require.NoError(t, err)
	proposal, found = app.CustomGovKeeper.GetProposal(ctx, proposalID)
	require.True(t, found)
	require.Equal(t, proposal.VotingEndTime.Unix(), ctx.BlockTime().Unix()+int64(properties.MinimumProposalEndTime))

	// check longer duration proposal
	app.CustomGovKeeper.SetProposalDuration(ctx, tsukitypes.ProposalTypeSetProposalDurations, 2400)
	proposalID, err = app.CustomGovKeeper.CreateAndSaveProposalWithContent(ctx, "title", "description", &types.SetProposalDurationsProposal{})
	require.NoError(t, err)
	proposal, found = app.CustomGovKeeper.GetProposal(ctx, proposalID)
	require.True(t, found)
	require.Equal(t, proposal.VotingEndTime.Unix(), ctx.BlockTime().Unix()+2400)
}
