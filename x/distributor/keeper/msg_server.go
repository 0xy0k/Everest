package keeper

import (
	"github.com/TsukiCore/tsuki/x/distributor/types"
)

type msgServer struct {
	keeper Keeper
	cgk    types.CustomGovKeeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, cgk types.CustomGovKeeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
		cgk:    cgk,
	}
}

var _ types.MsgServer = msgServer{}
