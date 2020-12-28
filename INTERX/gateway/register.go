package gateway

import (
	"github.com/TsukiCore/tsuki/INTERX/gateway/cosmos"
	"github.com/TsukiCore/tsuki/INTERX/gateway/interx"
	"github.com/TsukiCore/tsuki/INTERX/gateway/tsuki"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterRequest is a function to register requests.
func RegisterRequest(router *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	cosmos.RegisterRequest(router, gwCosmosmux, rpcAddr)
	tsuki.RegisterRequest(router, gwCosmosmux, rpcAddr)
	interx.RegisterRequest(router, gwCosmosmux, rpcAddr)
}
