package gov

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/TsukiCore/tsuki/x/gov/keeper"
	"github.com/TsukiCore/tsuki/x/gov/types"
	cumstomtypes "github.com/TsukiCore/tsuki/x/staking/types"
)

type Querier struct {
	keeper keeper.Keeper
}

func NewQuerier(keeper keeper.Keeper) types.QueryServer {
	return &Querier{keeper: keeper}
}

func (q Querier) RolesByAddress(ctx context.Context, request *types.RolesByAddressRequest) (*types.RolesByAddressResponse, error) {
	actor, found := q.keeper.GetNetworkActorByAddress(sdk.UnwrapSDKContext(ctx), request.ValAddr)
	if !found {
		return nil, cumstomtypes.ErrNetworkActorNotFound
	}

	return &types.RolesByAddressResponse{
		Roles: actor.Roles,
	}, nil
}

func (q Querier) CouncilorByAddress(ctx context.Context, request *types.CouncilorByAddressRequest) (*types.CouncilorResponse, error) {
	councilor, found := q.keeper.GetCouncilor(sdk.UnwrapSDKContext(ctx), request.ValAddr)
	if !found {
		return nil, types.ErrCouncilorNotFound
	}

	return &types.CouncilorResponse{Councilor: councilor}, nil
}

func (q Querier) CouncilorByMoniker(ctx context.Context, request *types.CouncilorByMonikerRequest) (*types.CouncilorResponse, error) {
	councilor, found := q.keeper.GetCouncilorByMoniker(sdk.UnwrapSDKContext(ctx), request.Moniker)
	if !found {
		return nil, types.ErrCouncilorNotFound
	}

	return &types.CouncilorResponse{Councilor: councilor}, nil
}

func (q Querier) PermissionsByAddress(ctx context.Context, request *types.PermissionsByAddressRequest) (*types.PermissionsResponse, error) {
	sdkContext := sdk.UnwrapSDKContext(ctx)

	networkActor, found := q.keeper.GetNetworkActorByAddress(sdkContext, request.ValAddr)
	if !found {
		return nil, cumstomtypes.ErrNetworkActorNotFound
	}

	return &types.PermissionsResponse{Permissions: networkActor.Permissions}, nil
}

func (q Querier) GetNetworkProperties(ctx context.Context, request *types.NetworkPropertiesRequest) (*types.NetworkPropertiesResponse, error) {
	sdkContext := sdk.UnwrapSDKContext(ctx)

	networkProperties := q.keeper.GetNetworkProperties(sdkContext)
	return &types.NetworkPropertiesResponse{Properties: networkProperties}, nil
}

func (q Querier) RolePermissions(ctx context.Context, request *types.RolePermissionsRequest) (*types.RolePermissionsResponse, error) {
	sdkContext := sdk.UnwrapSDKContext(ctx)

	perms, found := q.keeper.GetPermissionsForRole(sdkContext, types.Role(request.Role))
	if !found {
		return nil, types.ErrRoleDoesNotExist
	}

	return &types.RolePermissionsResponse{Permissions: &perms}, nil
}

func (q Querier) GetExecutionFee(ctx context.Context, request *types.ExecutionFeeRequest) (*types.ExecutionFeeResponse, error) {
	sdkContext := sdk.UnwrapSDKContext(ctx)
	fee := q.keeper.GetExecutionFee(sdkContext, request.TransactionType)
	if fee == nil {
		return nil, fmt.Errorf("fee does not exist for %s", request.TransactionType)
	}
	return &types.ExecutionFeeResponse{Fee: fee}, nil
}
