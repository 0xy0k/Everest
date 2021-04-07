package keeper

import (
	"context"
	"fmt"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	govkeeper "github.com/TsukiCore/tsuki/x/gov/keeper"
	customgovtypes "github.com/TsukiCore/tsuki/x/gov/types"
	"github.com/TsukiCore/tsuki/x/staking/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper    Keeper
	govKeeper govkeeper.Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, govKeeper govkeeper.Keeper) types.MsgServer {
	return &msgServer{
		keeper:    keeper,
		govKeeper: govKeeper,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) ClaimValidator(goCtx context.Context, msg *types.MsgClaimValidator) (*types.MsgClaimValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed := govkeeper.CheckIfAllowedPermission(ctx, k.govKeeper, sdk.AccAddress(msg.ValKey), customgovtypes.PermClaimValidator)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, "PermClaimValidator")
	}

	pk, ok := msg.PubKey.GetCachedValue().(cryptotypes.PubKey)
	if !ok {
		return nil, errors.Wrapf(errors.ErrInvalidPubKey, "Expecting cryptotypes.PubKey, got %T", pk)
	}

	_, err := k.keeper.GetValidator(ctx, msg.ValKey)
	if err == nil {
		return nil, types.ErrValidatorAlreadyClaimed
	}

	validator, err := types.NewValidator(msg.Moniker, msg.Website, msg.Social, msg.Identity, msg.Commission, msg.ValKey, pk)
	if err != nil {
		return nil, err
	}

	k.keeper.AddPendingValidator(ctx, validator)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeClaimValidator,
			sdk.NewAttribute(types.AttributeKeyMoniker, msg.Moniker),
			sdk.NewAttribute(types.AttributeKeyWebsite, msg.Website),
			sdk.NewAttribute(types.AttributeKeySocial, msg.Social),
			sdk.NewAttribute(types.AttributeKeyIdentity, msg.Identity),
			sdk.NewAttribute(types.AttributeKeyCommission, msg.Commission.String()),
			sdk.NewAttribute(types.AttributeKeyValKey, msg.ValKey.String()),
			sdk.NewAttribute(types.AttributeKeyPubKey, msg.PubKey.String()),
		),
	)
	return &types.MsgClaimValidatorResponse{}, nil
}

func (k msgServer) ProposalUnjailValidator(goCtx context.Context, msg *types.MsgProposalUnjailValidator) (*types.MsgProposalUnjailValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	isAllowed := govkeeper.CheckIfAllowedPermission(ctx, k.govKeeper, msg.Proposer, customgovtypes.PermCreateUnjailValidatorProposal)
	if !isAllowed {
		return nil, errors.Wrap(customgovtypes.ErrNotEnoughPermissions, customgovtypes.PermCreateUnjailValidatorProposal.String())
	}

	validator, err := k.keeper.GetValidatorByAccAddress(ctx, msg.Proposer)
	if err != nil {
		return nil, err
	}

	if !validator.IsJailed() {
		return nil, fmt.Errorf("validator is not jailed")
	}

	networkProperties := k.govKeeper.GetNetworkProperties(ctx)
	maxUnjailingTime := networkProperties.JailMaxTime

	info, found := k.keeper.GetValidatorJailInfo(ctx, validator.ValKey)
	if !found {
		return nil, fmt.Errorf("validator jailing info not found")
	}

	if info.Time.Add(time.Duration(maxUnjailingTime) * time.Minute).Before(ctx.BlockTime()) {
		return nil, fmt.Errorf("time to unjail passed")
	}

	proposalID, err := k.CreateAndSaveProposalWithContent(
		ctx,
		msg.Description,
		types.NewProposalUnjailValidator(
			msg.Proposer,
			msg.Hash,
			msg.Reference,
		),
	)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			customgovtypes.EventTypeSubmitProposal,
			sdk.NewAttribute(customgovtypes.AttributeKeyProposalId, fmt.Sprintf("%d", proposalID)),
			sdk.NewAttribute(customgovtypes.AttributeKeyProposalType, msg.Type()),
			sdk.NewAttribute(types.AttributeKeyDescription, msg.Description),
			sdk.NewAttribute(types.AttributeKeyHash, msg.Hash),
			sdk.NewAttribute(types.AttributeKeyReference, msg.Reference),
		),
	)
	return &types.MsgProposalUnjailValidatorResponse{
		ProposalID: proposalID,
	}, nil
}

func (k msgServer) CreateAndSaveProposalWithContent(ctx sdk.Context, description string, content customgovtypes.Content) (uint64, error) {
	blockTime := ctx.BlockTime()
	proposalID, err := k.govKeeper.GetNextProposalID(ctx)
	if err != nil {
		return 0, err
	}

	properties := k.govKeeper.GetNetworkProperties(ctx)

	proposal, err := customgovtypes.NewProposal(
		proposalID,
		content,
		blockTime,
		blockTime.Add(time.Second*time.Duration(properties.ProposalEndTime)),
		blockTime.Add(time.Second*time.Duration(properties.ProposalEndTime)+
			time.Second*time.Duration(properties.ProposalEnactmentTime),
		),
		ctx.BlockHeight()+2,
		ctx.BlockHeight()+3,
		description,
	)

	k.govKeeper.SaveProposal(ctx, proposal)
	k.govKeeper.AddToActiveProposals(ctx, proposal)

	return proposalID, nil
}
