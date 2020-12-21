package gateway

import (
	"net/http"

	functions "github.com/TsukiCore/tsuki/INTERX/functions"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

const (
	queryTsukiFunctions = "/api/tsuki/metadata"
)

// RegisterTsukiQueryRoutes registers tx query routers.
func RegisterTsukiQueryRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(queryTsukiFunctions, QueryTsukiFunctions(rpcAddr)).Methods(GET)

	AddRPCMethod(GET, queryTsukiFunctions, "This is an API to query tsuki functions and metadata.", true)
}

func queryTsukiFunctionsHandle(rpcAddr string) (interface{}, interface{}, int) {
	functions := functions.GetTsukiFunctions()

	return functions, nil, http.StatusOK
}

// QueryTsukiFunctions is a function to list functions and metadata.
func QueryTsukiFunctions(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := GetInterxRequest(r)
		response := GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		response.Response, response.Error, statusCode = queryTsukiFunctionsHandle(rpcAddr)

		WrapResponse(w, request, *response, statusCode, false)
	}
}
