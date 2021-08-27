package dto

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterMerchantRequest struct {
	Username          string `json:"username"`
	Password          string `json:"password"`
	MerchantName      string `json:"merchant_name"`
	HearOfficeAddress string `json:"head_office_address"`
}

type RegisterCustomerRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	CustomerName string `json:"customer_name"`
	Phone        string `json:"phone"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type RegisterMerchantResponse struct {
	UserID            int64  `json:"user_id"`
	Username          string `json:"username"`
	MerchantID        int64  `json:"merchant_id"`
	MerchantName      string `json:"merchant_name"`
	HearOfficeAddress string `json:"head_office_address"`
}

type RegisterCustomerResponse struct {
	UserID       int64  `json:"user_id"`
	Username     string `json:"username"`
	CustomerID   int64  `json:"customer_id"`
	CustomerName string `json:"customer_name"`
	Phone        string `json:"phone"`
}

func (r LoginRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Username, validation.Required); err != nil {
		return errs.NewBadRequestError("Username is required")

	}

	if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")

	}

	return nil
}

func NewRegisterUserMerchantResponse(data *domain.UserMerchant) *RegisterMerchantResponse {
	result := RegisterMerchantResponse{
		UserID:            data.UserID,
		Username:          data.Username,
		MerchantID:        data.MerchantID,
		MerchantName:      data.MerchantName,
		HearOfficeAddress: data.HearOfficeAddress,
	}

	return &result
}

func NewRegisterUserCustomerResponse(data *domain.UserCustomer) *RegisterCustomerResponse {
	result := RegisterCustomerResponse{
		UserID:       data.UserID,
		Username:     data.Username,
		CustomerID:   data.CustomerID,
		CustomerName: data.CustomerName,
		Phone:        data.Phone,
	}

	return &result
}

func (r RegisterMerchantRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Username, validation.Required); err != nil {
		return errs.NewBadRequestError("Username is required")

	}

	if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")

	}

	if err := validation.Validate(r.MerchantName, validation.Required); err != nil {
		return errs.NewBadRequestError("MerchantName is required")

	}

	if err := validation.Validate(r.HearOfficeAddress, validation.Required); err != nil {
		return errs.NewBadRequestError("HearOfficeAddress is required")

	}

	return nil
}

func (r RegisterCustomerRequest) Validate() *errs.AppError {

	if err := validation.Validate(r.Username, validation.Required); err != nil {
		return errs.NewBadRequestError("Username is required")

	}

	if err := validation.Validate(r.Password, validation.Required); err != nil {
		return errs.NewBadRequestError("Password is required")

	}

	if err := validation.Validate(r.CustomerName, validation.Required); err != nil {
		return errs.NewBadRequestError("CustomerName is required")

	}

	if err := validation.Validate(r.Phone, validation.Required); err != nil {
		return errs.NewBadRequestError("Phone is required")

	}

	return nil
}
