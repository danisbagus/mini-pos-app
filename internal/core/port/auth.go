package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IAuthRepo interface {
	FindOne(username string) (*domain.User, *errs.AppError)
	Verify(token string) *errs.AppError
	CreateUserMerchant(data *domain.UserMerchant) (*domain.UserMerchant, *errs.AppError)
	CreateUserCustomer(data *domain.UserCustomer) (*domain.UserCustomer, *errs.AppError)
}

type IAuthService interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
	RegisterMerchant(req *dto.RegisterMerchantRequest) (*dto.RegisterMerchantResponse, *errs.AppError)
	RegisterCustomer(req *dto.RegisterCustomerRequest) (*dto.RegisterCustomerResponse, *errs.AppError)
}
