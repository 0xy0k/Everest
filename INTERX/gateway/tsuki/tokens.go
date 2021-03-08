package tsuki

import (
	"net/http"

	"github.com/TsukiCore/tsuki/INTERX/common"
	"github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterTsukiTokensRoutes registers tsuki tokens query routers.
func RegisterTsukiTokensRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryTsukiTokensAliases, QueryTsukiTokensAliasesRequest(gwCosmosmux, rpcAddr)).Methods("GET")
	r.HandleFunc(config.QueryTsukiTokensRates, QueryTsukiTokensRatesRequest(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryTsukiTokensAliases, "This is an API to query all tokens aliases.", true)
	common.AddRPCMethod("GET", config.QueryTsukiTokensRates, "This is an API to query all tokens rates.", true)
}

func queryTsukiTokensAliasesHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryTsukiTokensAliasesRequest is a function to query all tokens aliases.
func QueryTsukiTokensAliasesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-tokens-aliases] Entering proposals query")

		if !common.RPCMethods["GET"][config.QueryTsukiTokensAliases].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryTsukiTokensAliases].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-tokens-aliases] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryTsukiTokensAliasesHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryTsukiTokensAliases].CachingEnabled)
	}
}

func queryTsukiTokensRatesHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryTsukiTokensRatesRequest is a function to query all tokens rates.
func QueryTsukiTokensRatesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-tokens-rates] Entering proposals query")

		if !common.RPCMethods["GET"][config.QueryTsukiTokensRates].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryTsukiTokensRates].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-tokens-rates] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryTsukiTokensRatesHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryTsukiTokensRates].CachingEnabled)
	}
}
