package staking

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/staking/keeper"
	"github.com/TsukiCore/tsuki/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyUnjailValidatorProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUnjailValidatorProposalHandler(keeper keeper.Keeper) *ApplyUnjailValidatorProposalHandler {
	return &ApplyUnjailValidatorProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyUnjailValidatorProposalHandler) ProposalType() string {
	return types.ProposalTypeUnjailValidator
}

func (a ApplyUnjailValidatorProposalHandler) Apply(ctx sdk.Context, proposal govtypes.Content) {
	p := proposal.(*types.ProposalUnjailValidator)

	err := a.keeper.Unjail(ctx, sdk.ValAddress(p.Proposer))
	if err != nil {
		panic("error unjailing")
	}
}
