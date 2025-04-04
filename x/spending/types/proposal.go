package types

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.Content = &UpdateSpendingPoolProposal{}

func NewUpdateSpendingPoolProposal(
	name string,
	claimStart uint64,
	claimEnd uint64,
	rates sdk.DecCoins,
	voteQuorum uint64,
	votePeriod uint64,
	voteEnactment uint64,
	owners PermInfo,
	beneficiaries WeightedPermInfo,
	dynamicRate bool,
	dynamicRatePeriod uint64,
) *UpdateSpendingPoolProposal {
	return &UpdateSpendingPoolProposal{
		Name:              name,
		ClaimStart:        claimStart,
		ClaimEnd:          claimEnd,
		Rates:             rates,
		VoteQuorum:        voteQuorum,
		VotePeriod:        votePeriod,
		VoteEnactment:     voteEnactment,
		Owners:            owners,
		Beneficiaries:     beneficiaries,
		DynamicRate:       dynamicRate,
		DynamicRatePeriod: dynamicRatePeriod,
	}
}

func (m *UpdateSpendingPoolProposal) ProposalType() string {
	return tsukitypes.ProposalTypeUpdateSpendingPool
}

func (m *UpdateSpendingPoolProposal) ProposalPermission() types.PermValue {
	return types.PermZero
}

func (m *UpdateSpendingPoolProposal) VotePermission() types.PermValue {
	return types.PermZero
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
	return types.PermZero
}

func (m *SpendingPoolDistributionProposal) VotePermission() types.PermValue {
	return types.PermZero
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
	return types.PermZero
}

func (m *SpendingPoolWithdrawProposal) VotePermission() types.PermValue {
	return types.PermZero
}

// ValidateBasic returns basic validation
func (m *SpendingPoolWithdrawProposal) ValidateBasic() error {
	return nil
}
