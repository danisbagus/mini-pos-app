package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type SaleTransactionHandler struct {
	Service port.ISaleTransactionService
}

func (rc SaleTransactionHandler) NewTransaction(w http.ResponseWriter, r *http.Request) {
	var request dto.NewSaleTransactionRequest

	claimData, err := GetClaimData(r)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	if claimData.Role != "CUSTOMER" {
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
