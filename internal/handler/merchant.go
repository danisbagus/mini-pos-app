package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/gorilla/mux"
)

type MerchantHandler struct {
	Service port.IMerchantService
}

func (rc MerchantHandler) GetMerchantList(w http.ResponseWriter, r *http.Request) {
	claimData, err := GetClaimData(r)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	if claimData.Role != "ADMIN" {
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

func (rc MerchantHandler) GetMerchantDetail(w http.ResponseWriter, r *http.Request) {
	claimData, err := GetClaimData(r)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	if claimData.Role != "ADMIN" {
		err := errs.NewAuthorizationError(fmt.Sprintf("%s role is not authorized", claimData.Role))
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	vars := mux.Vars(r)
	MerchantID, _ := strconv.Atoi(vars["merchant_id"])

	dataList, err := rc.Service.GetOne(int64(MerchantID))
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, dataList)
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

func (rc MerchantHandler) GetMerchantProductList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	MerchantID, _ := strconv.Atoi(vars["merchant_id"])

	dataList, err := rc.Service.GetAllMerchantProduct(int64(MerchantID))
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, dataList)
}

func (rc MerchantHandler) UpdateMerchant(w http.ResponseWriter, r *http.Request) {
	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	if claimData.Role != "ADMIN" {
		err := errs.NewAuthorizationError(fmt.Sprintf("%s role is not authorized", claimData.Role))
		writeResponse(w, err.Code, err.AsMessage())
		return

	}

	var request dto.UpdateMerchanteRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	vars := mux.Vars(r)
	MerchantID, _ := strconv.Atoi(vars["merchant_id"])

	err := rc.Service.UpdateMerchantByID(int64(MerchantID), &request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, map[string]bool{
		"success": true,
	})
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

	err := rc.Service.UpdateMerchantByUserID(claimData.UserID, &request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, map[string]bool{
		"success": true,
	})
}

func (rc MerchantHandler) RemoveMerchant(w http.ResponseWriter, r *http.Request) {
	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	if claimData.Role != "ADMIN" {
		err := errs.NewAuthorizationError(fmt.Sprintf("%s role is not authorized", claimData.Role))
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	vars := mux.Vars(r)
	MerchantID, _ := strconv.Atoi(vars["merchant_id"])

	err := rc.Service.RemoveMerchant(int64(MerchantID))
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, map[string]bool{
		"success": true,
	})
}
