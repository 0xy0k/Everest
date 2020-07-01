package createOrder

import sdk "github.com/TsukiCore/cosmos-sdk/types"

func HandleMessage(context sdk.Context, keeper Keeper, message Message) (*sdk.Result, error) {
	keeper.CreateOrder(context, message.OrderBookID, message.OrderType, message.Amount, message.LimitPrice, message.Curator)
	return &sdk.Result{}, nil
}