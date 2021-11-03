package types

import (
	govtypes "github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CustomGovKeeper defines the expected interface contract the tokens module requires
type CustomGovKeeper interface {
	CheckIfAllowedPermission(ctx sdk.Context, addr sdk.AccAddress, permValue govtypes.PermValue) bool
	GetNextProposalIDAndIncrement(ctx sdk.Context) uint64
	GetNetworkProperties(ctx sdk.Context) *govtypes.NetworkProperties
	SaveProposal(ctx sdk.Context, proposal govtypes.Proposal)
	AddToActiveProposals(ctx sdk.Context, proposal govtypes.Proposal)
	CreateAndSaveProposalWithContent(ctx sdk.Context, title, description string, content govtypes.Content) (uint64, error)
}

type CustomStakingKeeper interface {
	PauseProposalNotApprovedValidators(ctx sdk.Context, proposalID uint64) error
}
