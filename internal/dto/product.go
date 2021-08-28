package dto

import (
	"mime/multipart"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type NewProductRequest struct {
	ProductName string `json:"-"`
	UserID      int64  `json:"-"`
	MerchantID  int64  `json:"-"`
	Image       string `json:"-"`
	Quantity    int64  `json:"-"`
	Price       int64  `json:"-"`
	File        multipart.File
}

type NewProductResponse struct {
	SKUID       string `json:"sku_id"`
	MerchantID  int64  `json:"merchant_id"`
	ProductName string `json:"product_name"`
	Image       string `json:"image"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
}

func NewNewProductResponse(data *domain.ProductPrice) *NewProductResponse {
	result := &NewProductResponse{
		SKUID:       data.SKUID,
		MerchantID:  data.MerchantID,
		ProductName: data.ProductName,
		Image:       data.Image,
		Quantity:    data.Quantity,
		Price:       data.Price,
	}
	return result
}

func (r NewProductRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.ProductName, validation.Required); err != nil {
		return errs.NewBadRequestError("Product name is required")

	}

	if err := validation.Validate(r.Quantity, validation.Required); err != nil {
		return errs.NewBadRequestError("Product quantity is required")

	}

	if err := validation.Validate(r.Price, validation.Required); err != nil {
		return errs.NewBadRequestError("Product price is required")

	}

	if r.Quantity <= 0 {
		return errs.NewValidationError("Product quantity must more than 0")
	}

	if r.Price <= 0 {
		return errs.NewValidationError("Product price must more than 0")
	}
	return nil
}
