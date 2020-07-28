package listOrders

import (
	"fmt"
	"net/http"

	"github.com/TsukiCore/cosmos-sdk/client"
	"github.com/TsukiCore/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

func GetOrders(cliCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		max := vars["max_orders"]
		min := vars["min_amount"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/tsukiHub/listOrders/%s/%s/%s", id, max, min), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
