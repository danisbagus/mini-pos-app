package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/gorilla/mux"
)

type OutletHandler struct {
	Service port.IOutletService
}

func (rc OutletHandler) GetOutletListByMerchantID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	merchantID, _ := strconv.Atoi(vars["merchant_id"])

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

	dataList, err := rc.Service.GetAllByMerchantID(int64(merchantID))
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, dataList)
}
