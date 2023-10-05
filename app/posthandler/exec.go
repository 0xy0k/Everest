package posthandler

import (
	tsukitypes "github.com/TsukiCore/tsuki/types"
	feeprocessingkeeper "github.com/TsukiCore/tsuki/x/feeprocessing/keeper"
	customgovkeeper "github.com/TsukiCore/tsuki/x/gov/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// ExecutionDecorator stores execution status
type ExecutionDecorator struct {
	customGovKeeper     customgovkeeper.Keeper
	feeprocessingKeeper feeprocessingkeeper.Keeper
}

func NewExecutionDecorator(
	customGovKeeper customgovkeeper.Keeper,
	feeprocessingKeeper feeprocessingkeeper.Keeper,
) sdk.PostDecorator {
	return &ExecutionDecorator{
		customGovKeeper:     customGovKeeper,
		feeprocessingKeeper: feeprocessingKeeper,
	}
}

func (d ExecutionDecorator) PostHandle(ctx sdk.Context, tx sdk.Tx, simulate, success bool, next sdk.PostHandler) (sdk.Context, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	// execution fee should be prepaid
	for _, msg := range feeTx.GetMsgs() {
		fee := d.customGovKeeper.GetExecutionFee(ctx, tsukitypes.MsgType(msg))
		if fee == nil {
			return next(ctx, tx, simulate, success)
		}

		d.feeprocessingKeeper.SetExecutionStatusSuccess(ctx, msg)
	}

	return next(ctx, tx, simulate, success)
}
