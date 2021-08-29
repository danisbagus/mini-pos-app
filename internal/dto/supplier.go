package dto

import "github.com/danisbagus/mini-pos-app/internal/core/domain"

type SupplierResponse struct {
	SupplierID   int64  `json:"outlet_id"`
	SupplierName string `json:"outlet_name"`
}

type SupplierListResponse struct {
	Suppliers []SupplierResponse `json:"data"`
}

func NewGetListSupplierResponse(data []domain.Supplier) *SupplierListResponse {
	dataList := make([]SupplierResponse, len(data))

	for k, v := range data {
		dataList[k] = SupplierResponse{
			SupplierID:   v.SupplierID,
			SupplierName: v.SupplierName,
		}
	}
	return &SupplierListResponse{Suppliers: dataList}
}

func NewGetDetailSupplierResponse(data *domain.Supplier) *SupplierResponse {
	result := &SupplierResponse{
		SupplierID:   data.SupplierID,
		SupplierName: data.SupplierName,
	}
	return result
}
