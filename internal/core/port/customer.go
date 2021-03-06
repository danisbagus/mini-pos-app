package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type ICustomerRepo interface {
	FindAll() ([]domain.Customer, *errs.AppError)
	FindOne(customerID int64) (*domain.Customer, *errs.AppError)
	FindOneByUserID(userID int64) (*domain.UserCustomer, *errs.AppError)
	Update(customerID int64, data *domain.Customer) *errs.AppError
	Delete(customerID int64) *errs.AppError
}

type ICustomerService interface {
	GetAll() (*dto.CustomerListResponse, *errs.AppError)
	GetOne(customerID int64) (*dto.CustomerResponse, *errs.AppError)
	GetDetailByUserID(userID int64) (*dto.UserCustomerResponse, *errs.AppError)
	UpdateCustomerByUserID(userID int64, data *dto.UpdateCustomerRequest) *errs.AppError
	RemoveCustomer(customerID int64) *errs.AppError
}
