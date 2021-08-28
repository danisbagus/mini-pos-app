package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Service port.IProducService
}

func (rc ProductHandler) NewProduct(w http.ResponseWriter, r *http.Request) {

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

	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	fileLocation := fmt.Sprintf("public/uploads/%v%v", time.Now().UnixNano(), filepath.Ext(handler.Filename))

	quantity, err := strconv.Atoi(r.FormValue("quantity"))
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request := dto.NewProductRequest{
		UserID:      claimData.UserID,
		ProductName: r.FormValue("product_name"),
		Quantity:    int64(quantity),
		Image:       fileLocation,
		Price:       int64(price),
		File:        file,
	}

	data, appErr := rc.Service.NewProduct(&request)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	writeResponse(w, http.StatusCreated, data)
}

func (rc ProductHandler) GetProductDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	SKUID := vars["sku_id"]

	data, err := rc.Service.GetDetail(SKUID)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, data)
}
