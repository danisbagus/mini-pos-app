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

type CustomerHandler struct {
	Service port.ICustomerService
}

func (rc CustomerHandler) GetCustomerList(w http.ResponseWriter, r *http.Request) {
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

func (rc CustomerHandler) GetCustomerDetail(w http.ResponseWriter, r *http.Request) {
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
	CustomerID, _ := strconv.Atoi(vars["customer_id"])

	dataList, err := rc.Service.GetOne(int64(CustomerID))
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, dataList)
}

func (rc CustomerHandler) GetCustomerDetailMe(w http.ResponseWriter, r *http.Request) {
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
	var userID int64 = claimData.UserID

	data, err := rc.Service.GetDetailByUserID(userID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, data)
}

func (rc CustomerHandler) UpdateCustomerMe(w http.ResponseWriter, r *http.Request) {
	claimData, appErr := GetClaimData(r)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	if claimData.Role != "CUSTOMER" {
		err := errs.NewAuthorizationError(fmt.Sprintf("%s role is not authorized", claimData.Role))
		writeResponse(w, err.Code, err.AsMessage())
		return

	}

	var request dto.UpdateCustomerRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err := rc.Service.UpdateCustomerByUserID(claimData.UserID, &request)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, map[string]bool{
		"success": true,
	})
}
