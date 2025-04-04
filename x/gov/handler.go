package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/TsukiCore/tsuki/x/gov/keeper"
	"github.com/TsukiCore/tsuki/x/gov/types"
)

func NewHandler(ck keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(ck)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgSetNetworkProperties:
			res, err := msgServer.SetNetworkProperties(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgSetExecutionFee:
			res, err := msgServer.SetExecutionFee(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		// Permission Related
		case *types.MsgWhitelistPermissions:
			res, err := msgServer.WhitelistPermissions(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveWhitelistedPermissions:
			res, err := msgServer.RemoveWhitelistedPermissions(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBlacklistPermissions:
			res, err := msgServer.BlacklistPermissions(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveBlacklistedPermissions:
			res, err := msgServer.RemoveBlacklistedPermissions(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		// Councilor Related
		case *types.MsgClaimCouncilor:
			res, err := msgServer.ClaimCouncilor(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		// Role Related
		case *types.MsgWhitelistRolePermission:
			res, err := msgServer.WhitelistRolePermission(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgBlacklistRolePermission:
			res, err := msgServer.BlacklistRolePermission(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveWhitelistRolePermission:
			res, err := msgServer.RemoveWhitelistRolePermission(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRemoveBlacklistRolePermission:
			res, err := msgServer.RemoveBlacklistRolePermission(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCreateRole:
			res, err := msgServer.CreateRole(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgAssignRole:
			res, err := msgServer.AssignRole(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgUnassignRole:
			res, err := msgServer.UnassignRole(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		// Proposal related
		case *types.MsgSubmitProposal:
			res, err := msgServer.SubmitProposal(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgVoteProposal:
			res, err := msgServer.VoteProposal(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		// Poll Proposal related
		case *types.MsgPollCreate:
			res, err := msgServer.PollCreate(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgPollVote:
			res, err := msgServer.PollVote(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		// identity registrar related
		case *types.MsgRegisterIdentityRecords:
			res, err := msgServer.RegisterIdentityRecords(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgDeleteIdentityRecords:
			res, err := msgServer.DeleteIdentityRecords(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgRequestIdentityRecordsVerify:
			res, err := msgServer.RequestIdentityRecordsVerify(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgHandleIdentityRecordsVerifyRequest:
			res, err := msgServer.HandleIdentityRecordsVerifyRequest(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		case *types.MsgCancelIdentityRecordsVerifyRequest:
			res, err := msgServer.CancelIdentityRecordsVerifyRequest(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)
		default:
			return nil, errors.Wrapf(errors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}
	}
}
