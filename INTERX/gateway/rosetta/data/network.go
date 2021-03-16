package data

import (
	"encoding/json"
	"net/http"

	"github.com/TsukiCore/tsuki/INTERX/common"
	"github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/TsukiCore/tsuki/INTERX/types"
	"github.com/TsukiCore/tsuki/INTERX/types/rosetta"
	"github.com/TsukiCore/tsuki/INTERX/types/rosetta/dataapi"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// RegisterNetworkRoutes registers network API routers.
func RegisterNetworkRoutes(r *mux.Router, gwCosmosmux *runtime.ServeMux, rpcAddr string) {
	r.HandleFunc(config.QueryRosettaNetworkList, QueryNetworkListRequest(gwCosmosmux, rpcAddr)).Methods("POST")
	r.HandleFunc(config.QueryRosettaNetworkOptions, QueryNetworkOptionsRequest(gwCosmosmux, rpcAddr)).Methods("POST")

	common.AddRPCMethod("POST", config.QueryRosettaNetworkList, "This is an API to query network list.", true)
	common.AddRPCMethod("POST", config.QueryRosettaNetworkOptions, "This is an API to query network list.", true)
}

func queryNetworkListHandler(r *http.Request, request types.InterxRequest, rpcAddr string, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	var req dataapi.NetworkListRequest

	err := json.Unmarshal(request.Params, &req)
	if err != nil {
		common.GetLogger().Error("[rosetta-query-networklist] Failed to decode the request: ", err)
		return common.RosettaServeError(0, "failed to unmarshal", err.Error(), http.StatusBadRequest)
	}

	var response dataapi.NetworkListResponse

	success, failure, status := common.MakeGetRequest(rpcAddr, "/status", "")

	if success != nil {
		type TempResponse struct {
			NodeInfo struct {
				Network string `json:"network"`
			} `json:"node_info"`
		}
		result := TempResponse{}

		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[rosetta-query-networklist] Invalid response format", err)
			return common.RosettaServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		err = json.Unmarshal(byteData, &result)
		if err != nil {
			common.GetLogger().Error("[rosetta-query-networklist] Invalid response format", err)
			return common.RosettaServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		response.NetworkIdentifiers = make([]rosetta.NetworkIdentifier, 0)
		response.NetworkIdentifiers = append(response.NetworkIdentifiers, rosetta.NetworkIdentifier{
			Blockchain: "tsukid",
			Network:    result.NodeInfo.Network,
		})

		return response, nil, http.StatusOK
	}

	return nil, failure, status
}

// QueryNetworkListRequest is a function to query network list.
func QueryNetworkListRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[rosetta-query-networklist] Entering network list query")

		if !common.RPCMethods["POST"][config.QueryRosettaNetworkList].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["POST"][config.QueryRosettaNetworkList].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[rosetta-query-networklist] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryNetworkListHandler(r, request, rpcAddr, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["POST"][config.QueryRosettaNetworkList].CachingEnabled)
	}
}

func queryNetworkOptionsHandler(r *http.Request, request types.InterxRequest, rpcAddr string, gwCosmosmux *runtime.ServeMux) (interface{}, interface{}, int) {
	var req dataapi.NetworkOptionsRequest

	err := json.Unmarshal(request.Params, &req)
	if err != nil {
		common.GetLogger().Error("[rosetta-query-networkoptions] Failed to decode the request: ", err)
		return common.RosettaServeError(0, "failed to unmarshal", err.Error(), http.StatusBadRequest)
	}

	var response dataapi.NetworkOptionsResponse

	success, failure, status := common.MakeGetRequest(rpcAddr, "/status", "")

	if success != nil {
		type TempResponse struct {
			NodeInfo struct {
				Version string `json:"version"`
			} `json:"node_info"`
		}
		result := TempResponse{}

		byteData, err := json.Marshal(success)
		if err != nil {
			common.GetLogger().Error("[rosetta-query-networkoptions] Invalid response format", err)
			return common.RosettaServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		err = json.Unmarshal(byteData, &result)
		if err != nil {
			common.GetLogger().Error("[rosetta-query-networkoptions] Invalid response format", err)
			return common.RosettaServeError(0, "", err.Error(), http.StatusInternalServerError)
		}

		response.Version = rosetta.Version{
			NodeVersion:       result.NodeInfo.Version,
			MiddlewareVersion: config.InterxVersion,
		}

		return response, nil, http.StatusOK
	}

	return nil, failure, status
}

// QueryNetworkOptionsRequest is a function to query network options.
func QueryNetworkOptionsRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[rosetta-query-networkoptions] Entering network list query")

		if !common.RPCMethods["POST"][config.QueryRosettaNetworkOptions].Enabled {
			response.Response, response.Error, statusCode = common.ServeError(0, "", "API disabled", http.StatusForbidden)
		} else {
			if common.RPCMethods["POST"][config.QueryRosettaNetworkOptions].CachingEnabled {
				found, cacheResponse, cacheError, cacheStatus := common.SearchCache(request, response)
				if found {
					response.Response, response.Error, statusCode = cacheResponse, cacheError, cacheStatus
					common.WrapResponse(w, request, *response, statusCode, false)

					common.GetLogger().Info("[rosetta-query-networkoptions] Returning from the cache")
					return
				}
			}

			response.Response, response.Error, statusCode = queryNetworkOptionsHandler(r, request, rpcAddr, gwCosmosmux)
		}

		common.WrapResponse(w, request, *response, statusCode, common.RPCMethods["POST"][config.QueryRosettaNetworkList].CachingEnabled)
	}
}
