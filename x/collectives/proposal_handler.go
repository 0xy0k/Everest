package collectives

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/collectives/keeper"
	"github.com/TsukiCore/tsuki/x/collectives/types"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyCollectiveSendDonationProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCollectiveSendDonationProposalHandler(keeper keeper.Keeper) *ApplyCollectiveSendDonationProposalHandler {
	return &ApplyCollectiveSendDonationProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyCollectiveSendDonationProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeCollectiveSendDonation
}

func (a ApplyCollectiveSendDonationProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalCollectiveSendDonation)
	return a.keeper.SendDonation(ctx, p.Name, p.Address, p.Amount)
}

type ApplyCollectiveUpdateProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCollectiveUpdateProposalHandler(keeper keeper.Keeper) *ApplyCollectiveUpdateProposalHandler {
	return &ApplyCollectiveUpdateProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyCollectiveUpdateProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeCollectiveUpdate
}

func (a ApplyCollectiveUpdateProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalCollectiveUpdate)

	a.keeper.SetCollective(ctx, p.Collective)
	return nil
}

type ApplyCollectiveRemoveProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCollectiveRemoveProposalHandler(keeper keeper.Keeper) *ApplyCollectiveRemoveProposalHandler {
	return &ApplyCollectiveRemoveProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyCollectiveRemoveProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeCollectiveRemove
}

func (a ApplyCollectiveRemoveProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalCollectiveRemove)
	a.keeper.DeleteCollective(ctx, p.Name)
	return nil
}
