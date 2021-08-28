package dto

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserCustomerResponse struct {
	UserID       int64  `json:"user_id"`
	Role         string `json:"role"`
	Username     string `json:"username"`
	CustomerID   int64  `json:"Customer_id"`
	CustomerName string `json:"Customer_name"`
	Phone        string `json:"phone"`
}

type UpdateCustomerRequest struct {
	CustomerName string `json:"customer_name"`
	Phone        string `json:"phone"`
}

func NewGetDetailUserCustomerResponse(data *domain.UserCustomer) *UserCustomerResponse {
	result := &UserCustomerResponse{
		UserID:       data.UserID,
		Role:         data.Role,
		Username:     data.Username,
		CustomerID:   data.CustomerID,
		CustomerName: data.CustomerName,
		Phone:        data.Phone,
	}
	return result
}

func (r UpdateCustomerRequest) Validate() *errs.AppError {
	if err := validation.Validate(r.CustomerName, validation.Required); err != nil {
		return errs.NewBadRequestError("CustomerName is required")
	}

	if err := validation.Validate(r.Phone, validation.Required); err != nil {
		return errs.NewBadRequestError("Phone quantity is required")
	}
	return nil
}
