package dto

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type OutletResponse struct {
	OutletID   int64  `json:"outlet_id"`
	MerchantID int64  `json:"merchant_id"`
	OutletName string `json:"outlet_name"`
	Address    string `json:"address"`
}

type OutletListResponse struct {
	Outlets []OutletResponse `json:"data"`
}

type NewOutletRequest struct {
	OutletName string `json:"outlet_name"`
	Address    string `json:"address"`
	UserID     int64  `json:"-"`
}

type NewOutletResponse struct {
	OutletID   int64  `json:"outlet_id"`
	MerchantID int64  `json:"merchant_id"`
	OutletName string `json:"outlet_name"`
	Address    string `json:"address"`
}

func NewGetListOutletResponse(data []domain.Outlet) *OutletListResponse {
	dataList := make([]OutletResponse, len(data))

	for k, v := range data {
		dataList[k] = OutletResponse{
			OutletID:   v.OutletID,
			MerchantID: v.MerchantID,
			OutletName: v.OutletName,
			Address:    v.Address,
		}
	}
	return &OutletListResponse{Outlets: dataList}
}

func NewGetDetailOutletResponse(data *domain.Outlet) *OutletResponse {
	result := &OutletResponse{
		OutletID:   data.OutletID,
		MerchantID: data.MerchantID,
		OutletName: data.OutletName,
		Address:    data.Address,
	}
	return result
}

func NewNewOutletResponse(data *domain.Outlet) *NewOutletResponse {
	result := &NewOutletResponse{
		OutletID:   data.OutletID,
		MerchantID: data.MerchantID,
		OutletName: data.OutletName,
		Address:    data.Address,
	}

	return result
}

func (r NewOutletRequest) Validate() *errs.AppError {
	if err := validation.Validate(r.OutletName, validation.Required); err != nil {
		return errs.NewBadRequestError("OutletName is required")
	}
	if err := validation.Validate(r.Address, validation.Required); err != nil {
		return errs.NewBadRequestError("Address is required")
	}
	return nil
}
