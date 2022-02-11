package spending

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/spending/keeper"
	"github.com/TsukiCore/tsuki/x/spending/types"
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

func (a ApplyUpdateSpendingPoolProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content) error {
	p := proposal.(*spendingtypes.UpdateSpendingPoolProposal)

	pool := a.keeper.GetSpendingPool(ctx, p.Name)
	if pool == nil {
		return types.ErrPoolDoesNotExist
	}

	a.keeper.SetSpendingPool(ctx, types.SpendingPool{
		Name:          p.Name,
		ClaimStart:    p.ClaimStart,
		ClaimEnd:      p.ClaimEnd,
		Expire:        p.Expire,
		Token:         p.Token,
		Rate:          p.Rate,
		VoteQuorum:    p.VoteQuorum,
		VotePeriod:    p.VotePeriod,
		VoteEnactment: p.VoteEnactment,
		Owners:        &p.Owners,
		Beneficiaries: &p.Beneficiaries,
		Balance:       pool.Balance,
	})

	return nil
}

type ApplySpendingPoolDistributionProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySpendingPoolDistributionProposalHandler(keeper keeper.Keeper) *ApplySpendingPoolDistributionProposalHandler {
	return &ApplySpendingPoolDistributionProposalHandler{
		keeper: keeper,
	}
}

func (a ApplySpendingPoolDistributionProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeSpendingPoolDistribution
}

func (a ApplySpendingPoolDistributionProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content) error {
	p := proposal.(*spendingtypes.SpendingPoolDistributionProposal)
	_ = p
	// TODO: should distribute all the tokens to beneficiaries
	return nil
}

type ApplySpendingPoolWithdrawProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySpendingPoolWithdrawProposalHandler(keeper keeper.Keeper) *ApplySpendingPoolWithdrawProposalHandler {
	return &ApplySpendingPoolWithdrawProposalHandler{
		keeper: keeper,
	}
}

func (a ApplySpendingPoolWithdrawProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeSpendingPoolWithdraw
}

func (a ApplySpendingPoolWithdrawProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content) error {
	p := proposal.(*spendingtypes.SpendingPoolWithdrawProposal)
	_ = p
	// TODO: should withdraw specified amount of tokens to beneficiaries
	return nil
}
