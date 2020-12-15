package rest

import (
	"bytes"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/TsukiCore/tsuki/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
)

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/slashing/validators/{validatorAddr}/activate", NewActivateRequestHandlerFn(clientCtx)).Methods("POST")
}

// Activate TX body
type ActivateReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
}

// NewActivateRequestHandlerFn returns an HTTP REST handler for creating a MsgActivate
// transaction.
func NewActivateRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bech32Validator := vars["validatorAddr"]

		var req ActivateReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, err := sdk.AccAddressFromBech32(req.BaseReq.From)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		valAddr, err := sdk.ValAddressFromBech32(bech32Validator)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		if !bytes.Equal(fromAddr, valAddr) {
			rest.WriteErrorResponse(w, http.StatusUnauthorized, "must use own validator address")
			return
		}

		msg := types.NewMsgActivate(valAddr)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
