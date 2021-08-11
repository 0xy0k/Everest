package tsuki

import (
	"net/http"

	"github.com/TsukiCore/tsuki/INTERX/common"
	"github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterTsukiUpgradeRoutes registers tsuki upgrade query routers.
func RegisterTsukiUpgradeRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryUpgradePlan, QueryUpgradePlanRequest(gwCosmosmux, rpcAddr)).Methods("GET")

	common.AddRPCMethod("GET", config.QueryUpgradePlan, "This is an API to query upgrade plan.", true)
}

func queryUpgradePlanHandler(r *http.Request, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	return common.ServeGRPC(r, gwCosmosmux)
}

// QueryUpgradePlanRequest is a function to query current upgrade plan.
func QueryUpgradePlanRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-upgrade-plan] Entering upgrade plan query")

		if !common.RPCMethods["GET"][config.QueryUpgradePlan].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["GET"][config.QueryUpgradePlan].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[query-upgrade-plan] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryUpgradePlanHandler(r, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["GET"][config.QueryUpgradePlan].CachingEnabled)
	}
}
