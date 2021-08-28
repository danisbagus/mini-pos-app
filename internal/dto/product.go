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

type UpdateProductRequest struct {
	ProductName string `json:"-"`
	UserID      int64  `json:"-"`
	MerchantID  int64  `json:"-"`
	Image       string `json:"-"`
	Quantity    int64  `json:"-"`
	File        multipart.File
}

type UpdateProductResponse struct {
	SKUID       string `json:"sku_id"`
	MerchantID  int64  `json:"merchant_id"`
	ProductName string `json:"product_name"`
	Image       string `json:"image"`
	Quantity    int64  `json:"quantity"`
}

type UpdateProductPriceRequest struct {
	OutletID int64 `json:"outlet_id"`
	Price    int64 `json:"price"`
	UserID   int64 `json:"-"`
}

type ProductPrice struct {
	OutletID int64 `json:"outlet_id"`
	Price    int64 `json:"price"`
}

type ProductResponse struct {
	SKUID       string         `json:"sku_id"`
	MerchantID  int64          `json:"merchant_id"`
	ProductName string         `json:"product_name"`
	Image       string         `json:"image"`
	Quantity    int64          `json:"quantity"`
	Price       []ProductPrice `json:"prices"`
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

func NewUpdateProductResponse(data *domain.Product) *UpdateProductResponse {
	result := &UpdateProductResponse{
		SKUID:       data.SKUID,
		MerchantID:  data.MerchantID,
		ProductName: data.ProductName,
		Image:       data.Image,
		Quantity:    data.Quantity,
	}

	return result
}

func NewGetDetailProductResponse(data *domain.Product, prices []domain.Prices) *ProductResponse {
	productPrice := make([]ProductPrice, len(prices))

	for k, v := range prices {
		productPrice[k] = ProductPrice{
			OutletID: v.OutletID,
			Price:    v.Price,
		}
	}
	result := &ProductResponse{
		SKUID:       data.SKUID,
		MerchantID:  data.MerchantID,
		ProductName: data.ProductName,
		Image:       data.Image,
		Quantity:    data.Quantity,
		Price:       productPrice,
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

func (r UpdateProductRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.ProductName, validation.Required); err != nil {
		return errs.NewBadRequestError("Product name is required")

	}

	if err := validation.Validate(r.Quantity, validation.Required); err != nil {
		return errs.NewBadRequestError("Product quantity is required")

	}

	if r.Quantity <= 0 {
		return errs.NewValidationError("Product quantity must more than 0")
	}

	return nil
}

func (r UpdateProductPriceRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.OutletID, validation.Required); err != nil {
		return errs.NewBadRequestError("Outlite ID is required")

	}

	if err := validation.Validate(r.Price, validation.Required); err != nil {
		return errs.NewBadRequestError("Price ID is required")

	}

	if r.Price <= 0 {
		return errs.NewValidationError("Price must more than 0")
	}

	return nil
}
