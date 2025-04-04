package types

import (
	"errors"

	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ types.Content = &ProposalUpsertTokenAlias{}
	_ types.Content = &ProposalUpsertTokenRates{}
	_ types.Content = &ProposalTokensWhiteBlackChange{}
)

func NewUpsertTokenAliasProposal(
	symbol string,
	name string,
	icon string,
	decimals uint32,
	denoms []string,
	isInvalidated bool,
) *ProposalUpsertTokenAlias {
	return &ProposalUpsertTokenAlias{
		Symbol:      symbol,
		Name:        name,
		Icon:        icon,
		Decimals:    decimals,
		Denoms:      denoms,
		Invalidated: isInvalidated,
	}
}

func (m *ProposalUpsertTokenAlias) ProposalType() string {
	return tsukitypes.ProposalTypeUpsertTokenAlias
}

func (m *ProposalUpsertTokenAlias) ProposalPermission() types.PermValue {
	return types.PermCreateUpsertTokenAliasProposal
}

func (m *ProposalUpsertTokenAlias) VotePermission() types.PermValue {
	return types.PermVoteUpsertTokenAliasProposal
}

// ValidateBasic returns basic validation
func (m *ProposalUpsertTokenAlias) ValidateBasic() error {
	return nil
}

func NewUpsertTokenRatesProposal(
	denom string,
	rate sdk.Dec,
	feePayments bool,
	stakeCap sdk.Dec,
	stakeMin sdk.Int,
	stakeToken bool,
	isInvalidated bool,
) *ProposalUpsertTokenRates {
	return &ProposalUpsertTokenRates{
		Denom:       denom,
		Rate:        rate,
		FeePayments: feePayments,
		StakeCap:    stakeCap,
		StakeMin:    stakeMin,
		StakeToken:  stakeToken,
		Invalidated: isInvalidated,
	}
}

func (m *ProposalUpsertTokenRates) ProposalType() string {
	return tsukitypes.ProposalTypeUpsertTokenRates
}

func (m *ProposalUpsertTokenRates) ProposalPermission() types.PermValue {
	return types.PermCreateUpsertTokenRateProposal
}

func (m *ProposalUpsertTokenRates) VotePermission() types.PermValue {
	return types.PermVoteUpsertTokenRateProposal
}

// ValidateBasic returns basic validation
func (m *ProposalUpsertTokenRates) ValidateBasic() error {
	if m.StakeCap.LT(sdk.NewDec(0)) { // not positive
		return errors.New("reward cap should be positive")
	}

	if m.StakeCap.GT(sdk.OneDec()) { // more than 1
		return errors.New("reward cap not be more than 100%")
	}
	return nil
}

func NewTokensWhiteBlackChangeProposal(isBlacklist, isAdd bool, tokens []string) *ProposalTokensWhiteBlackChange {
	return &ProposalTokensWhiteBlackChange{isBlacklist, isAdd, tokens}
}

func (m *ProposalTokensWhiteBlackChange) ProposalType() string {
	return tsukitypes.ProposalTypeTokensWhiteBlackChange
}

func (m *ProposalTokensWhiteBlackChange) ProposalPermission() types.PermValue {
	return types.PermCreateTokensWhiteBlackChangeProposal
}

func (m *ProposalTokensWhiteBlackChange) VotePermission() types.PermValue {
	return types.PermVoteTokensWhiteBlackChangeProposal
}

// ValidateBasic returns basic validation
func (m *ProposalTokensWhiteBlackChange) ValidateBasic() error {
	return nil
}
