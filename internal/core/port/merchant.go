package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IMerchantRepo interface {
	FindAll() ([]domain.Merchant, *errs.AppError)
	FindOneByID(merchantID int64) (*domain.UserMerchant, *errs.AppError)
	FindOneByUserID(userID int64) (*domain.UserMerchant, *errs.AppError)
	Update(merchatID int64, data *domain.Merchant) *errs.AppError
	Delete(merchatID int64) *errs.AppError
}

type IMerchantService interface {
	GetAll() (*dto.MerchantListResponse, *errs.AppError)
	GetOne(merchantID int64) (*dto.MerchantResponse, *errs.AppError)
	GetDetailByUserID(userID int64) (*dto.UserMerchantResponse, *errs.AppError)
	GetAllMerchantProduct(merchantID int64) (*dto.ProductListResponse, *errs.AppError)
	UpdateMerchantByID(merchantID int64, data *dto.UpdateMerchanteRequest) *errs.AppError
	UpdateMerchantByUserID(userID int64, data *dto.UpdateMerchanteRequest) *errs.AppError
	RemoveMerchant(merchantID int64) *errs.AppError
}
