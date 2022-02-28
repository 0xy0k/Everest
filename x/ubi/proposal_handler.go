package ubi

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/ubi/keeper"
	ubitypes "github.com/TsukiCore/tsuki/x/ubi/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyUpsertUBIProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpsertUBIProposalHandler(keeper keeper.Keeper) *ApplyUpsertUBIProposalHandler {
	return &ApplyUpsertUBIProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyUpsertUBIProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeUpsertUBI
}

func (a ApplyUpsertUBIProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*ubitypes.UpsertUBIProposal)

	// TODO: The proposal should fail if sum of all ((float)amount / period) * 31556952 for all UBI records is greater than ubi-hard-cap.

	_ = p
	return nil
}

type ApplyRemoveUBIProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyRemoveUBIProposalHandler(keeper keeper.Keeper) *ApplyRemoveUBIProposalHandler {
	return &ApplyRemoveUBIProposalHandler{keeper: keeper}
}

func (a ApplyRemoveUBIProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeRemoveUBI
}

func (a ApplyRemoveUBIProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*ubitypes.RemoveUBIProposal)
	_ = p
	return nil
}
