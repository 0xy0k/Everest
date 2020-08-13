package tsukiHub

import (
	"strings"

	"github.com/TsukiCore/cosmos-sdk/client"
	"github.com/gorilla/mux"

	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
	"github.com/TsukiCore/tsuki/x/tsukiHub/queries/listOrderBooks"
	"github.com/TsukiCore/tsuki/x/tsukiHub/queries/listOrders"
)

func RegisterRESTRoutes(cliContext client.Context, router *mux.Router) {
	//router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.CreateOrderBookTransaction}, "/"), createOrderBook.RestRequestHandler(cliContext)).Methods("POST")
	//router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.CreateOrderTransaction}, "/"), createOrder.RestRequestHandler(cliContext)).Methods("POST")

	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.ListOrderBooksQuery}, "/"), listOrderBooks.GetOrderBooks(cliContext)).Methods("GET")
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.ListOrderBooksQueryByTP}, "/"), listOrderBooks.GetOrderBooksByTP(cliContext)).Methods("GET")
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.ListOrders}, "/"), listOrders.GetOrders(cliContext)).Methods("GET")
}
