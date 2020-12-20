package gateway

import (
	"net/http"

	functions "github.com/TsukiCore/tsuki/INTERX/functions"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

const (
	queryTsukiFunctions = "/api/tsuki/metadata"
	queryStatus        = "/api/cosmos/status"
)

// RegisterTsukiQueryRoutes registers tx query routers.
func RegisterTsukiQueryRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(queryTsukiFunctions, QueryTsukiFunctions(rpcAddr)).Methods(GET)
	r.HandleFunc(queryStatus, QueryStatusRequest(rpcAddr)).Methods(GET)

	AddRPCMethod(GET, queryTsukiFunctions, "This is an API to query tsuki functions and metadata.", true)
	AddRPCMethod(GET, queryStatus, "This is an API to query status.", true)
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

func queryStatusHandle(rpcAddr string) (interface{}, interface{}, int) {
	return MakeGetRequest(rpcAddr, "/status", "")
}

// QueryStatusRequest is a function to query status.
func QueryStatusRequest(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := GetInterxRequest(r)
		response := GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		if !rpcMethods[GET][queryStatus].Enabled {
			response.Response, response.Error, statusCode = ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if rpcMethods[GET][queryStatus].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					WrapResponse(w, request, *response, statusCode, false)
					return
				}
			}

			response.Response, response.Error, statusCode = queryStatusHandle(rpcAddr)
		}

		WrapResponse(w, request, *response, statusCode, rpcMethods[GET][queryStatus].CachingEnabled)
	}
}
