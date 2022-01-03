package tsuki

import (
	"fmt"
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
	type TokenAliasesResult struct {
		Decimals int64    `json:"decimals"`
		Denoms   []string `json:"denoms"`
		Name     string   `json:"name"`
		Symbol   string   `json:"symbol"`
		Icon     string   `json:"icon"`
		Amount   int64    `json:"amount,string"`
	}

	tokens := common.GetTokenAliases(gwCosmosmux, r.Clone(r.Context()))
	tokensSupply := common.GetTokenSupply(gwCosmosmux, r.Clone(r.Context()))

	fmt.Println(tokens, tokensSupply)

	result := make([]TokenAliasesResult, 0)

	for _, token := range tokens {
		flag := false
		for _, denom := range token.Denoms {
			for _, supply := range tokensSupply {
				if denom == supply.Denom {
					result = append(result, TokenAliasesResult{
						Decimals: token.Decimals,
						Denoms:   token.Denoms,
						Name:     token.Name,
						Symbol:   token.Symbol,
						Icon:     token.Icon,
						Amount:   supply.Amount,
					})

					flag = true
					break
				}
			}
			if flag {
				break
			}
		}
	}

	return result, nil, http.StatusOK
}

// QueryTsukiTokensAliasesRequest is a function to query all tokens aliases.
func QueryTsukiTokensAliasesRequest(gwCosmosmux *runtime.ServeMux, rpcAddr string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := common.GetInterxRequest(r)
		response := common.GetResponseFormat(request, rpcAddr)
		statusCode := http.StatusOK

		common.GetLogger().Info("[query-tokens-aliases] Entering token aliases query")

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

		common.GetLogger().Info("[query-tokens-rates] Entering token rates query")

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
