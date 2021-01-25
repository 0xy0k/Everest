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
	r.HandleFunc(common.QueryTsukiStatus, QueryTsukiStatusRequest(rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", common.QueryTsukiFunctions, "This is an API to query tsuki functions and metadata.", true)
	common.AddRPCMethod("GET", common.QueryTsukiStatus, "This is an API to query tsuki status.", true)
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

// QueryTsukiStatusRequest is a function to query tsuki status.
func QueryTsukiStatusRequest(rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-tsuki-status] Entering status query")

		if !common.RPCMethods["GET"][common.QueryTsukiStatus].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][common.QueryTsukiStatus].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-tsuki-status] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = common.MakeGetRequest(rpcAddr, "/status", "")
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][common.QueryTsukiStatus].CachingEnabled)
	}
}
