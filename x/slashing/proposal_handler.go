package slashing

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/slashing/keeper"
	"github.com/TsukiCore/tsuki/x/slashing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyResetWholeValidatorRankProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyResetWholeValidatorRankProposalHandler(keeper keeper.Keeper) *ApplyResetWholeValidatorRankProposalHandler {
	return &ApplyResetWholeValidatorRankProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyResetWholeValidatorRankProposalHandler) ProposalType() string {
	return types.ProposalTypeResetWholeValidatorRank
}

func (a ApplyResetWholeValidatorRankProposalHandler) Apply(ctx sdk.Context, proposal govtypes.Content) error {
	_ = proposal.(*types.ProposalResetWholeValidatorRank)

	return a.keeper.ResetWholeValidatorRank(ctx)
}
