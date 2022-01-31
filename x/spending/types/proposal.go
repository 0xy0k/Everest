package types

import (
	"time"

	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.Content = &UpdateSpendingPoolProposal{}

func NewUpdateSpendingPoolProposal(
	name string,
	claimStart time.Time,
	claimEnd time.Time,
	expire uint64,
	token string,
	rate sdk.Dec,
	voteQuorum uint64,
	votePeriod uint64,
	voteEnactment uint64,
	owners PermInfo,
	beneficiaries PermInfo,
) *UpdateSpendingPoolProposal {
	return &UpdateSpendingPoolProposal{
		Name:          name,
		ClaimStart:    claimStart,
		ClaimEnd:      claimEnd,
		Expire:        expire,
		Token:         token,
		Rate:          rate,
		VoteQuorum:    voteQuorum,
		VotePeriod:    votePeriod,
		VoteEnactment: voteEnactment,
		Owners:        owners,
		Beneficiaries: beneficiaries,
	}
}

func (m *UpdateSpendingPoolProposal) ProposalType() string {
	return tsukitypes.ProposalTypeUpdateSpendingPool
}

func (m *UpdateSpendingPoolProposal) ProposalPermission() types.PermValue {
	return types.PermCreateUpdateSpendingPoolProposal
}

func (m *UpdateSpendingPoolProposal) VotePermission() types.PermValue {
	return types.PermVoteUpdateSpendingPoolProposal
}

// ValidateBasic returns basic validation
func (m *UpdateSpendingPoolProposal) ValidateBasic() error {
	return nil
}

var _ types.Content = &SpendingPoolDistributionProposal{}

func NewSpendingPoolDistributionProposal(
	name string,
) *SpendingPoolDistributionProposal {
	return &SpendingPoolDistributionProposal{
		PoolName: name,
	}
}

func (m *SpendingPoolDistributionProposal) ProposalType() string {
	return tsukitypes.ProposalTypeSpendingPoolDistribution
}

func (m *SpendingPoolDistributionProposal) ProposalPermission() types.PermValue {
	return types.PermCreateSpendingPoolDistributionProposal
}

func (m *SpendingPoolDistributionProposal) VotePermission() types.PermValue {
	return types.PermVoteSpendingPoolDistributionProposal
}

// ValidateBasic returns basic validation
func (m *SpendingPoolDistributionProposal) ValidateBasic() error {
	return nil
}

var _ types.Content = &SpendingPoolWithdrawProposal{}

func NewSpendingPoolWithdrawProposal(
	name string,
	beneficiaries []string,
	amounts sdk.Coins,
) *SpendingPoolWithdrawProposal {
	return &SpendingPoolWithdrawProposal{
		PoolName:      name,
		Beneficiaries: beneficiaries,
		Amounts:       amounts,
	}
}

func (m *SpendingPoolWithdrawProposal) ProposalType() string {
	return tsukitypes.ProposalTypeSpendingPoolWithdraw
}

func (m *SpendingPoolWithdrawProposal) ProposalPermission() types.PermValue {
	return types.PermCreateSpendingPoolWithdrawProposal
}

func (m *SpendingPoolWithdrawProposal) VotePermission() types.PermValue {
	return types.PermVoteSpendingPoolWithdrawProposal
}

// ValidateBasic returns basic validation
func (m *SpendingPoolWithdrawProposal) ValidateBasic() error {
	return nil
}
