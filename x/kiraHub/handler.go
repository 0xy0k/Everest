package tsukiHub

import (
	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrderBook"
	"github.com/pkg/errors"

	sdk "github.com/TsukiCore/cosmos-sdk/types"
)


func NewHandler(keeper Keeper) sdk.Handler {
	return func(context sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		context = context.WithEventManager(sdk.NewEventManager())

		switch message := msg.(type) {
		case createOrderBook.Message:
			return createOrderBook.HandleMessage(context, keeper.getCreateOrderBookKeeper(), message)

		default:
			return nil, errors.Wrapf(constants.UnknownMessageCode, "%T", msg)
		}
	}
}

