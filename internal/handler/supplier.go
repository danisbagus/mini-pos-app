package handler

import (
	"fmt"
	"net/http"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type SuppplierHandler struct {
	Service port.ISupplierService
}

func (rc SuppplierHandler) GetSuppplierList(w http.ResponseWriter, r *http.Request) {

	claimData, err := GetClaimData(r)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	if !(claimData.Role == "MERCHANT" || claimData.Role == "ADMIN") {
		err := errs.NewAuthorizationError(fmt.Sprintf("%s role is not authorized", claimData.Role))
		writeResponse(w, err.Code, err.AsMessage())
		return

	}

	dataList, err := rc.Service.GetAll()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, dataList)
}
