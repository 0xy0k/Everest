package tokens

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/spending/keeper"
	spendingtypes "github.com/TsukiCore/tsuki/x/spending/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyUpdateSpendingPoolProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpdateSpendingPoolProposalHandler(keeper keeper.Keeper) *ApplyUpdateSpendingPoolProposalHandler {
	return &ApplyUpdateSpendingPoolProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyUpdateSpendingPoolProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeUpdateSpendingPool
}

func (a ApplyUpdateSpendingPoolProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*spendingtypes.ProposalUpsertTokenAlias)

	tokenAlians := spendingtypes.NewTokenAlias(p.Symbol, p.Name, p.Icon, p.Decimals, p.Denoms)
	return a.keeper.UpsertTokenAlias(ctx, *tokenAlians)
}

type ApplySpendingPoolDistributionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpdateSpendingPoolProposalHandler(keeper keeper.Keeper) *ApplySpendingPoolDistributionProposalHandler {
	return &ApplySpendingPoolDistributionProposalHandler{
		keeper: keeper,
	}
}

func (a ApplySpendingPoolDistributionProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeSpendingPoolDistribution
}

func (a ApplySpendingPoolDistributionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*spendingtypes.ProposalUpsertTokenAlias)

	tokenAlians := spendingtypes.NewTokenAlias(p.Symbol, p.Name, p.Icon, p.Decimals, p.Denoms)
	return a.keeper.UpsertTokenAlias(ctx, *tokenAlians)
}

type ApplySpendingPoolWithdrawProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpdateSpendingPoolProposalHandler(keeper keeper.Keeper) *ApplySpendingPoolWithdrawProposalHandler {
	return &ApplySpendingPoolWithdrawProposalHandler{
		keeper: keeper,
	}
}

func (a ApplySpendingPoolWithdrawProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeSpendingPoolWithdraw
}

func (a ApplySpendingPoolWithdrawProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*spendingtypes.ProposalUpsertTokenAlias)

	tokenAlians := spendingtypes.NewTokenAlias(p.Symbol, p.Name, p.Icon, p.Decimals, p.Denoms)
	return a.keeper.UpsertTokenAlias(ctx, *tokenAlians)
}
