package basket

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/basket/keeper"
	"github.com/TsukiCore/tsuki/x/basket/types"
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyCreateBasketProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCreateBasketProposalHandler(keeper keeper.Keeper) *ApplyCreateBasketProposalHandler {
	return &ApplyCreateBasketProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyCreateBasketProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeCreateBasket
}

func (a ApplyCreateBasketProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash sdk.Dec) error {
	p := proposal.(*types.ProposalCreateBasket)
	return a.keeper.CreateBasket(ctx, p.Basket)
}

type ApplyEditBasketProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyEditBasketProposalHandler(keeper keeper.Keeper) *ApplyEditBasketProposalHandler {
	return &ApplyEditBasketProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyEditBasketProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeEditBasket
}

func (a ApplyEditBasketProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash sdk.Dec) error {
	p := proposal.(*types.ProposalEditBasket)

	return a.keeper.EditBasket(ctx, p.Basket)
}

type ApplyBasketWithdrawSurplusProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyBasketWithdrawSurplusProposalHandler(keeper keeper.Keeper) *ApplyBasketWithdrawSurplusProposalHandler {
	return &ApplyBasketWithdrawSurplusProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyBasketWithdrawSurplusProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeBasketWithdrawSurplus
}

func (a ApplyBasketWithdrawSurplusProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash sdk.Dec) error {
	p := proposal.(*types.ProposalBasketWithdrawSurplus)
	return a.keeper.BasketWithdrawSurplus(ctx, *p)
}
