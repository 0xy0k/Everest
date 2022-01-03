package upgrade

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/upgrade/keeper"
	upgradetypes "github.com/TsukiCore/tsuki/x/upgrade/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplySoftwareUpgradeProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplySoftwareUpgradeProposalHandler(keeper keeper.Keeper) *ApplySoftwareUpgradeProposalHandler {
	return &ApplySoftwareUpgradeProposalHandler{
		keeper: keeper,
	}
}

func (a ApplySoftwareUpgradeProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeSoftwareUpgrade
}

func (a ApplySoftwareUpgradeProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	p := proposal.(*upgradetypes.ProposalSoftwareUpgrade)

	plan := upgradetypes.NewUpgradePlan(
		p.Name,
		p.Resources,
		p.UpgradeTime,
		p.OldChainId,
		p.NewChainId,
		p.MaxEnrolmentDuration,
		p.RollbackChecksum,
		p.InstateUpgrade,
		p.RebootRequired,
		p.SkipHandler,
		proposalID,
	)
	err := a.keeper.SaveNextPlan(ctx, plan)
	return err
}

type ApplyCancelSoftwareUpgradeProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyCancelSoftwareUpgradeProposalHandler(keeper keeper.Keeper) *ApplyCancelSoftwareUpgradeProposalHandler {
	return &ApplyCancelSoftwareUpgradeProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyCancelSoftwareUpgradeProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeCancelSoftwareUpgrade
}

func (a ApplyCancelSoftwareUpgradeProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content) error {
	a.keeper.ClearNextPlan(ctx)
	return nil
}
