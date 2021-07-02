package types

import (
	customgovtypes "github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CustomGovKeeper defines the expected interface contract the tokens module requires
type CustomGovKeeper interface {
	CheckIfAllowedPermission(ctx sdk.Context, addr sdk.AccAddress, permValue customgovtypes.PermValue) bool
	GetNextProposalID(ctx sdk.Context) uint64
	GetNetworkProperties(ctx sdk.Context) *customgovtypes.NetworkProperties
	SaveProposal(ctx sdk.Context, proposal customgovtypes.Proposal)
	AddToActiveProposals(ctx sdk.Context, proposal customgovtypes.Proposal)
}
