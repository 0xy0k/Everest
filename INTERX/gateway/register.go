package gateway

import (
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterRequest is a function to register requests.
func RegisterRequest(router *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	RegisterCosmosAuthRoutes(router, gwCosmosmux, rpcAddr)
	RegisterCosmosBankRoutes(router, gwCosmosmux, rpcAddr)
	RegisterCosmosTxRoutes(router, gwCosmosmux, rpcAddr)

	RegisterInterxFaucetRoutes(router, gwCosmosmux, rpcAddr)
	RegisterInterxQueryRoutes(router, gwCosmosmux, rpcAddr)
	RegisterInterxTxRoutes(router, gwCosmosmux, rpcAddr)

	RegisterTsukiGovRoutes(router, gwCosmosmux, rpcAddr)
	RegisterTsukiQueryRoutes(router, gwCosmosmux, rpcAddr)
}
