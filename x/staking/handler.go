package staking

import (
	customgovtypes "github.com/TsukiCore/tsuki/x/gov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	govkeeper "github.com/TsukiCore/tsuki/x/gov/keeper"
	customkeeper "github.com/TsukiCore/tsuki/x/staking/keeper"
	"github.com/TsukiCore/tsuki/x/staking/types"
)

func NewHandler(ck customkeeper.Keeper, govkeeper govkeeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgClaimValidator:
			return handleMsgClaimValidator(ctx, ck, govkeeper, msg)
		default:
			return nil, errors.Wrapf(errors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}

func handleMsgClaimValidator(ctx sdk.Context, k customkeeper.Keeper, govkeeper govkeeper.Keeper, msg *types.MsgClaimValidator) (*sdk.Result, error) {
	valPubKey, err := sdk.GetPubKeyFromBech32(sdk.Bech32PubKeyTypeConsPub, msg.PubKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get consensus node public key")
	}

	addr := sdk.AccAddress(msg.ValKey)
	networkActor, err := govkeeper.GetNetworkActorByAddress(ctx, addr)
	if err != nil {
		return nil, types.ErrNetworkActorNotFound
	}

	if !networkActor.HasRole(customgovtypes.RoleValidator) {
		if networkActor.Permissions.IsBlacklisted(customgovtypes.PermClaimValidator) {
			return nil, errors.Wrap(types.ErrNotEnoughPermissions, "PermClaimValidator is blacklisted")
		}

		if !networkActor.Permissions.IsWhitelisted(customgovtypes.PermClaimValidator) {
			return nil, errors.Wrap(types.ErrNotEnoughPermissions, "PermClaimValidator not whitelisted")
		}
	}

	validator, err := types.NewValidator(msg.Moniker, msg.Website, msg.Social, msg.Identity, msg.Commission, msg.ValKey, valPubKey)
	if err != nil {
		return nil, err
	}

	k.AddValidator(ctx, validator)

	return &sdk.Result{}, nil
}
