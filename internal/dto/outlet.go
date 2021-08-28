package dto

import "github.com/danisbagus/mini-pos-app/internal/core/domain"

type OutletResponse struct {
	OutletID   int64  `json:"outlet_id"`
	MerchantID int64  `json:"merchant_id"`
	OutletName string `json:"outlet_name"`
	Address    string `json:"address"`
}

type OutletListResponse struct {
	Outlets []OutletResponse `json:"data"`
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
