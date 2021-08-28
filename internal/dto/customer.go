package dto

import "github.com/danisbagus/mini-pos-app/internal/core/domain"

type UserCustomerResponse struct {
	UserID       int64  `json:"user_id"`
	Role         string `json:"role"`
	Username     string `json:"username"`
	CustomerID   int64  `json:"Customer_id"`
	CustomerName string `json:"Customer_name"`
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
