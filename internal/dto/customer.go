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

type CustomerResponse struct {
	CustomerID   int64  `json:"customer_id"`
	UserID       int64  `json:"user_id"`
	CustomerName string `json:"customer_name"`
	Phone        string `json:"phone"`
}

type CustomerListResponse struct {
	Customer []CustomerResponse `json:"data"`
}

func NewGetListCustomerResponse(data []domain.Customer) *CustomerListResponse {
	dataList := make([]CustomerResponse, len(data))

	for k, v := range data {
		dataList[k] = CustomerResponse{
			CustomerID:   v.CustomerID,
			CustomerName: v.CustomerName,
			UserID:       v.UserID,
			Phone:        v.Phone,
		}
	}
	return &CustomerListResponse{Customer: dataList}
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

func NewGetDetailCustomerResponse(data *domain.Customer) *CustomerResponse {
	result := &CustomerResponse{
		UserID:       data.UserID,
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
