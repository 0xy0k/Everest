package tsukiHub

import (
	"github.com/TsukiCore/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
	"strings"

	constants "github.com/TsukiCore/tsuki/x/tsukiHub/constants"
	"github.com/TsukiCore/tsuki/x/tsukiHub/queries/listOrderBooks"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrder"
	"github.com/TsukiCore/tsuki/x/tsukiHub/transactions/createOrderBook"
)

func RegisterRESTRoutes(cliContext context.CLIContext, router *mux.Router) {
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.CreateOrderBookTransaction}, "/"), createOrderBook.RestRequestHandler(cliContext)).Methods("POST")
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.CreateOrderTransaction}, "/"), createOrder.RestRequestHandler(cliContext)).Methods("POST")

	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.ListOrderBooksQuery}, "/"), listOrderBooks.GetOrderBooks(cliContext)).Methods("GET")
	router.HandleFunc(strings.Join([]string{"", TransactionRoute, constants.ListOrderBooksQueryByTP}, "/"), listOrderBooks.GetOrderBooksByTP(cliContext)).Methods("GET")
}