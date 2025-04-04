package tokens

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	"github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/tokens/keeper"
	tokenstypes "github.com/TsukiCore/tsuki/x/tokens/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ApplyUpsertTokenAliasProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpsertTokenAliasProposalHandler(keeper keeper.Keeper) *ApplyUpsertTokenAliasProposalHandler {
	return &ApplyUpsertTokenAliasProposalHandler{
		keeper: keeper,
	}
}

func (a ApplyUpsertTokenAliasProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeUpsertTokenAlias
}

func (a ApplyUpsertTokenAliasProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content, slash sdk.Dec) error {
	p := proposal.(*tokenstypes.ProposalUpsertTokenAlias)

	alias := tokenstypes.NewTokenAlias(p.Symbol, p.Name, p.Icon, p.Decimals, p.Denoms, p.Invalidated)
	return a.keeper.UpsertTokenAlias(ctx, *alias)
}

type ApplyUpsertTokenRatesProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyUpsertTokenRatesProposalHandler(keeper keeper.Keeper) *ApplyUpsertTokenRatesProposalHandler {
	return &ApplyUpsertTokenRatesProposalHandler{keeper: keeper}
}

func (a ApplyUpsertTokenRatesProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeUpsertTokenRates
}

func (a ApplyUpsertTokenRatesProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content, slash sdk.Dec) error {
	p := proposal.(*tokenstypes.ProposalUpsertTokenRates)

	rate := tokenstypes.NewTokenRate(p.Denom, p.Rate, p.FeePayments, p.StakeCap, p.StakeMin, p.StakeToken, p.Invalidated)
	return a.keeper.UpsertTokenRate(ctx, *rate)
}

type ApplyWhiteBlackChangeProposalHandler struct {
	keeper keeper.Keeper
}

func NewApplyWhiteBlackChangeProposalHandler(keeper keeper.Keeper) *ApplyWhiteBlackChangeProposalHandler {
	return &ApplyWhiteBlackChangeProposalHandler{keeper: keeper}
}

func (a ApplyWhiteBlackChangeProposalHandler) ProposalType() string {
	return tsukitypes.ProposalTypeTokensWhiteBlackChange
}

func (a ApplyWhiteBlackChangeProposalHandler) Apply(ctx sdk.Context, proposalID uint64, proposal types.Content, slash sdk.Dec) error {
	p := proposal.(*tokenstypes.ProposalTokensWhiteBlackChange)

	if p.IsBlacklist {
		if p.IsAdd {
			a.keeper.AddTokensToBlacklist(ctx, p.Tokens)
		} else {
			a.keeper.RemoveTokensFromBlacklist(ctx, p.Tokens)
		}
	} else {
		if p.IsAdd {
			a.keeper.AddTokensToWhitelist(ctx, p.Tokens)
		} else {
			a.keeper.RemoveTokensFromWhitelist(ctx, p.Tokens)
		}
	}
	return nil
}
