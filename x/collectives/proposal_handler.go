package collectives

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/collectives/keeper"
	"github.com/TsukiCore/tsuki/x/collectives/types"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyCreatecollectivesProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCreatecollectivesProposalHandler(keeper keeper.Keeper) *ApplyCreatecollectivesProposalHandler {
	return &ApplyCreatecollectivesProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyCreatecollectivesProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeCreateCollectives
}

func (a ApplyCreatecollectivesProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalCreateCollectives)
	return a.keeper.Createcollectives(ctx, p.collectives)
}

type ApplyEditcollectivesProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyEditcollectivesProposalHandler(keeper keeper.Keeper) *ApplyEditcollectivesProposalHandler {
	return &ApplyEditcollectivesProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyEditcollectivesProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeEditcollectives
}

func (a ApplyEditcollectivesProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalEditcollectives)

	return a.keeper.Editcollectives(ctx, p.collectives)
}

type ApplycollectivesWithdrawSurplusProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplycollectivesWithdrawSurplusProposalHandler(keeper keeper.Keeper) *ApplycollectivesWithdrawSurplusProposalHandler {
	return &ApplycollectivesWithdrawSurplusProposalHandler{
		keeper: keeper,
	}
}

func (a ApplycollectivesWithdrawSurplusProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypecollectivesWithdrawSurplus
}

func (a ApplycollectivesWithdrawSurplusProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalcollectivesWithdrawSurplus)
	return a.keeper.collectivesWithdrawSurplus(ctx, *p)
}
