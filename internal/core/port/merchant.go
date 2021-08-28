package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IMerchantRepo interface {
	FindOneByID(merchantID int64) (*domain.UserMerchant, *errs.AppError)
	FindOneByUserID(userID int64) (*domain.UserMerchant, *errs.AppError)
	Update(merchatID int64, data *domain.Merchant) *errs.AppError
}

type IMerchantService interface {
	GetDetailByUserID(userID int64) (*dto.UserMerchantResponse, *errs.AppError)
	UpdateMerchantByUserID(userID int64, data *dto.UpdateMerchanteRequest) *errs.AppError
}
