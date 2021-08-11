package tsuki

import (
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterRequest is a function to register requests.
func RegisterRequest(router *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	RegisterTsukiGovRoutes(router, gwCosmosmux, rpcAddr)
	RegisterTsukiGovProposalRoutes(router, gwCosmosmux, rpcAddr)
	RegisterTsukiQueryRoutes(router, gwCosmosmux, rpcAddr)
	RegisterTsukiTokensRoutes(router, gwCosmosmux, rpcAddr)
	RegisterTsukiUpgradeRoutes(router, gwCosmosmux, rpcAddr)
}
