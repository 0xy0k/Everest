package types

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewProposalCollectiveSendDonation(
	name string,
	address string,
	amounts sdk.Coins,
) *ProposalCollectiveSendDonation {
	return &ProposalCollectiveSendDonation{
		Name:    name,
		Address: address,
		Amounts: amounts,
	}
}

func (m *ProposalCollectiveSendDonation) ProposalType() string {
	return tsukitypes.ProposalTypeCollectiveSendDonation
}

func (m *ProposalCollectiveSendDonation) ProposalPermission() types.PermValue {
	return types.PermZero
}

func (m *ProposalCollectiveSendDonation) VotePermission() types.PermValue {
	return types.PermZero
}

// ValidateBasic returns basic validation
func (m *ProposalCollectiveSendDonation) ValidateBasic() error {
	return nil
}

func NewProposalCollectiveUpdate(
	name, description string,
	status CollectiveStatus,
	depositWhitelist DepositWhitelist,
	ownersWhitelist OwnersWhitelist,
	weightedSpendingPool []WeightedSpendingPool,
	claimStart, claimPeriod, claimEnd,
	voteQuorum, votePeriod, voteEnactment uint64,
) *ProposalCollectiveUpdate {
	return &ProposalCollectiveUpdate{
		Name:             name,
		Description:      description,
		Status:           status,
		DepositWhitelist: depositWhitelist,
		OwnersWhitelist:  ownersWhitelist,
		SpendingPools:    weightedSpendingPool,
		ClaimStart:       claimStart,
		ClaimPeriod:      claimPeriod,
		ClaimEnd:         claimEnd,
		VoteQuorum:       voteQuorum,
		VotePeriod:       votePeriod,
		VoteEnactment:    voteEnactment,
	}
}

func (m *ProposalCollectiveUpdate) ProposalType() string {
	return tsukitypes.ProposalTypeCollectiveUpdate
}

func (m *ProposalCollectiveUpdate) ProposalPermission() types.PermValue {
	return types.PermZero
}

func (m *ProposalCollectiveUpdate) VotePermission() types.PermValue {
	return types.PermZero
}

// ValidateBasic returns basic validation
func (m *ProposalCollectiveUpdate) ValidateBasic() error {
	return nil
}

func NewProposalCollectiveRemove(name string) *ProposalCollectiveRemove {
	return &ProposalCollectiveRemove{
		Name: name,
	}
}

func (m *ProposalCollectiveRemove) ProposalType() string {
	return tsukitypes.ProposalTypeCollectiveRemove
}

func (m *ProposalCollectiveRemove) ProposalPermission() types.PermValue {
	return types.PermZero
}

func (m *ProposalCollectiveRemove) VotePermission() types.PermValue {
	return types.PermZero
}

// ValidateBasic returns basic validation
func (m *ProposalCollectiveRemove) ValidateBasic() error {
	return nil
}
