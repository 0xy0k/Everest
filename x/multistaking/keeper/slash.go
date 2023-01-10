package keeper

import (
	"github.com/TsukiCore/tsuki/x/multistaking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func (k Keeper) SlashStakingPool(ctx sdk.Context, validator string, slash sdk.Dec) {
	pool, found := k.GetStakingPoolByValidator(ctx, validator)
	if !found {
		return
	}
	pool.Slashed = slash
	pool.Enabled = false

	totalStakingTokens := sdk.Coins{}
	for _, stakingToken := range pool.TotalStakingTokens {
		totalStakingTokens = totalStakingTokens.Add(sdk.NewCoin(stakingToken.Denom, stakingToken.Amount.ToDec().Mul(sdk.OneDec().Sub(pool.Slashed)).RoundInt()))
	}
	totalSlashedTokens := sdk.Coins(pool.TotalStakingTokens).Sub(totalStakingTokens)
	pool.TotalStakingTokens = totalStakingTokens

	bondDenom := k.govKeeper.BondDenom(ctx)
	bondDenomAmount := totalSlashedTokens.AmountOf(bondDenom)
	burnAmount := sdk.Coins{sdk.NewCoin(bondDenom, bondDenomAmount)}
	err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, burnAmount)
	if err != nil {
		panic(err)
	}

	treasurySendAmount := totalSlashedTokens.Sub(burnAmount)
	err = k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.ModuleName, authtypes.FeeCollectorName, treasurySendAmount)
	if err != nil {
		panic(err)
	}
	feesTreasury := k.distrKeeper.GetFeesTreasury(ctx)
	feesTreasury = feesTreasury.Add(treasurySendAmount...)
	k.distrKeeper.SetFeesTreasury(ctx, feesTreasury)
	k.SetStakingPool(ctx, pool)
}
