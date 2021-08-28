package dto

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type NewSaleTransactionRequest struct {
	UserID   int64  `json:"-"`
	SKUID    string `json:"sku_id"`
	OutletID int64  `json:"outlet_id"`
	Quantity int64  `json:"quantity"`
}

type NewSaleTransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	CustomerID    int64  `json:"customer_id"`
	SKUID         string `json:"sku_id"`
	OutletID      int64  `json:"OutletID"`
	Quantity      int64  `json:"quantity"`
	TotalPrice    int64  `json:"total_price"`
	CreatedAt     string `json:"created_at"`
}

type SaleTransactionList struct {
	SaleTransaction []NewSaleTransactionResponse `json:"data"`
}

func NewGetSaleTransactionReport(data []domain.SaleTransaction) *SaleTransactionList {
	dataList := make([]NewSaleTransactionResponse, len(data))

	for k, v := range data {
		dataList[k] = NewSaleTransactionResponse{
			TransactionID: v.TransactionID,
			CustomerID:    v.CustomerID,
			SKUID:         v.SKUID,
			OutletID:      v.OutletID,
			Quantity:      v.Quantity,
			TotalPrice:    v.TotalPrice,
			CreatedAt:     v.CreatedAt,
		}
	}
	return &SaleTransactionList{SaleTransaction: dataList}
}

func NewNewSaleTransactionResponse(data *domain.SaleTransaction) *NewSaleTransactionResponse {
	result := &NewSaleTransactionResponse{
		TransactionID: data.TransactionID,
		CustomerID:    data.CustomerID,
		SKUID:         data.SKUID,
		OutletID:      data.OutletID,
		Quantity:      data.Quantity,
		TotalPrice:    data.TotalPrice,
		CreatedAt:     data.CreatedAt,
	}

	return result
}

func (r NewSaleTransactionRequest) Validate() *errs.AppError {
	if err := validation.Validate(r.SKUID, validation.Required); err != nil {
		return errs.NewBadRequestError("SKU ID is required")
	}

	if err := validation.Validate(r.OutletID, validation.Required); err != nil {
		return errs.NewBadRequestError("OutletID quantity is required")
	}

	if err := validation.Validate(r.Quantity, validation.Required); err != nil {
		return errs.NewBadRequestError("Quantity price is required")
	}

	if r.Quantity <= 0 {
		return errs.NewValidationError("Quantity must more than 0")
	}

	return nil
}
