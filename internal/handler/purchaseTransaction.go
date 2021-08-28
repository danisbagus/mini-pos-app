package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type PurchaseTransactionHandler struct {
	Service port.IPurchaseTransactionService
}

func (rc PurchaseTransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.NewPurchaseTransactionRequest

	claimData, err := GetClaimData(r)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	if claimData.Role != "MERCHANT" {
		err := errs.NewAuthorizationError(fmt.Sprintf("%s role is not authorized", claimData.Role))
		writeResponse(w, err.Code, err.AsMessage())
		return

	}
	request.UserID = claimData.UserID

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	data, err := rc.Service.NewTransaction(&request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusCreated, data)
}

func (rc PurchaseTransactionHandler) GetTransactionReport(w http.ResponseWriter, r *http.Request) {
	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	if claimData.Role != "MERCHANT" {
		err := errs.NewAuthorizationError(fmt.Sprintf("%s role is not authorized", claimData.Role))
		writeResponse(w, err.Code, err.AsMessage())
		return

	}
	dataList, err := rc.Service.GetTransactionReport(claimData.UserID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, dataList)
}
