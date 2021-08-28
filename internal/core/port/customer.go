package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type ICustomerRepo interface {
	FindAll() ([]domain.Customer, *errs.AppError)
	// FindOne() (*domain.Customer, *errs.AppError)
	FindOneByUserID(userID int64) (*domain.UserCustomer, *errs.AppError)
	Update(customerID int64, data *domain.Customer) *errs.AppError
}

type ICustomerService interface {
	GetAll() (*dto.CustomerListResponse, *errs.AppError)
	GetDetailByUserID(userID int64) (*dto.UserCustomerResponse, *errs.AppError)
	UpdateCustomerByUserID(userID int64, data *dto.UpdateCustomerRequest) *errs.AppError
}
