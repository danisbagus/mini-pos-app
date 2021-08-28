package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type MerchantHandler struct {
	Service port.IMerchantService
}

func (rc MerchantHandler) GetMerchantDetailMe(w http.ResponseWriter, r *http.Request) {
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
	var userID int64 = claimData.UserID

	data, err := rc.Service.GetDetailByUserID(userID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, data)
}

func (rc MerchantHandler) UpdateMerchantMe(w http.ResponseWriter, r *http.Request) {
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

	var request dto.UpdateMerchanteRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := rc.Service.UpdateProductByUserID(claimData.UserID, &request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, map[string]bool{
		"success": true,
	})
}
