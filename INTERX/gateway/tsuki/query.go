package tsuki

import (
	"net/http"

	"github.com/TsukiCore/tsuki/INTERX/common"
	functions "github.com/TsukiCore/tsuki/INTERX/functions"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterTsukiQueryRoutes registers tx query routers.
func RegisterTsukiQueryRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(common.QueryTsukiFunctions, QueryTsukiFunctions(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", common.QueryTsukiFunctions, "This is an API to query tsuki functions and metadata.", true)
}

func queryTsukiFunctionsHandle(rpcAddr string) (interface{}, interface{}, int) {
	functions := functions.GetTsukiFunctions()

	return functions, nil, http.StatusOK
}

// QueryTsukiFunctions is a function to list functions and metadata.
func QueryTsukiFunctions(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		response.Response, response.Error, statusCode = queryTsukiFunctionsHandle(rpcAddr)

		common.WrapResponse(w, request, *response, statusCode, false)
	}
}
