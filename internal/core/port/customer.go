package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type ICustomerRepo interface {
	FindOneByUserID(userID int64) (*domain.UserCustomer, *errs.AppError)
}

type ICustomerService interface {
	GetDetailByUserID(userID int64) (*dto.UserCustomerResponse, *errs.AppError)
}
