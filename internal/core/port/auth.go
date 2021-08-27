package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IAuthRepo interface {
	FindOne(username string) (*domain.User, *errs.AppError)
}

type IAuthService interface {
	Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}
