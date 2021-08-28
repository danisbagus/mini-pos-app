package dto

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type NewPurchaseTransactionRequest struct {
	UserID     int64  `json:"-"`
	SKUID      string `json:"sku_id"`
	SuppierID  int64  `json:"supplier_id"`
	Quantity   int64  `json:"quantity"`
	TotalPrice int64  `json:"total_price"`
}

type NewPurchaseTransactionResponse struct {
	TransactionID string `json:"transaction_id"`
	MerchantID    int64  `json:"merchant_id"`
	SKUID         string `json:"sku_id"`
	SuppierID     int64  `json:"supplier_id"`
	Quantity      int64  `json:"quantity"`
	TotalPrice    int64  `json:"total_price"`
	CreatedAt     string `json:"created_at"`
}

func NewNewPurchaseTransactionResponse(data *domain.PurchaseTransaction) *NewPurchaseTransactionResponse {
	result := &NewPurchaseTransactionResponse{
		TransactionID: data.TransactionID,
		MerchantID:    data.MerchantID,
		SKUID:         data.SKUID,
		SuppierID:     data.SuppierID,
		Quantity:      data.Quantity,
		TotalPrice:    data.TotalPrice,
		CreatedAt:     data.CreatedAt,
	}

	return result
}

type PurchaseTransactionList struct {
	PurchaseTransaction []NewPurchaseTransactionResponse `json:"data"`
}

func NewGetPurchaseTransactionReport(data []domain.PurchaseTransaction) *PurchaseTransactionList {
	dataList := make([]NewPurchaseTransactionResponse, len(data))

	for k, v := range data {
		dataList[k] = NewPurchaseTransactionResponse{
			TransactionID: v.TransactionID,
			MerchantID:    v.MerchantID,
			SKUID:         v.SKUID,
			SuppierID:     v.SuppierID,
			Quantity:      v.Quantity,
			TotalPrice:    v.TotalPrice,
			CreatedAt:     v.CreatedAt,
		}
	}
	return &PurchaseTransactionList{PurchaseTransaction: dataList}
}

func (r NewPurchaseTransactionRequest) Validate() *errs.AppError {
	if err := validation.Validate(r.SKUID, validation.Required); err != nil {
		return errs.NewBadRequestError("SKU ID is required")
	}

	if err := validation.Validate(r.SuppierID, validation.Required); err != nil {
		return errs.NewBadRequestError("SuppierID quantity is required")
	}

	if err := validation.Validate(r.Quantity, validation.Required); err != nil {
		return errs.NewBadRequestError("Quantity price is required")
	}

	if err := validation.Validate(r.TotalPrice, validation.Required); err != nil {
		return errs.NewBadRequestError("TotalPrice price is required")
	}

	if r.Quantity <= 0 {
		return errs.NewValidationError("Quantity must more than 0")
	}

	if r.TotalPrice <= 0 {
		return errs.NewValidationError("TotalPrice price must more than 0")
	}

	return nil
}
