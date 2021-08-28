package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IMerchantRepo interface {
	FindOneByUserID(userID int64) (*domain.UserMerchant, *errs.AppError)
}

type IMerchantService interface {
	GetDetaiByUserID(userID int64) (*dto.UserMerchantResponse, *errs.AppError)
}