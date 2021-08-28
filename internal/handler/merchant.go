package handler

import (
	"net/http"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
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

	var userID int64 = claimData.UserID

	data, err := rc.Service.GetDetaiByUserID(userID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, data)
}
