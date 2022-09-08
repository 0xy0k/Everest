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

func (a ApplyCreateBasketProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalCreateBasket)

	a.keeper.SetBasket(ctx, p.Basket)
	return nil
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

func (a ApplyEditBasketProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalEditBasket)

	_, err := a.keeper.GetBasketById(ctx, p.Basket.Id)
	if err != nil {
		return err
	}

	a.keeper.SetBasket(ctx, p.Basket)
	return nil
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

func (a ApplyBasketWithdrawSurplusProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal govtypes.Content, slash uint64) error {
	p := proposal.(*types.ProposalBasketWithdrawSurplus)

	_, err := a.keeper.GetBasketById(ctx, p.BasketId)
	if err != nil {
		return err
	}

	// TODO: implement
	// a.keeper.SetBasket(ctx, p.Basket)
	return nil
}
