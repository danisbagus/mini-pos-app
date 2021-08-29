package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IOutletRepo interface {
	Create(data *domain.Outlet) (*domain.Outlet, *errs.AppError)
	FindAllByMerchantID(merchantID int64) ([]domain.Outlet, *errs.AppError)
	FindOneByID(outletID int64) (*domain.Outlet, *errs.AppError)
	Update(outletID int64, data *domain.Outlet) *errs.AppError
	Delete(outletID int64) *errs.AppError
}

type IOutletService interface {
	NewOutlet(data *dto.NewOutletRequest) (*dto.NewOutletResponse, *errs.AppError)
	GetAllByMerchantID(merchantID int64) (*dto.OutletListResponse, *errs.AppError)
	UpdateOutlet(outletID int64, data *dto.NewOutletRequest) *errs.AppError
	RemoveOutlet(outletID int64, userID int64) *errs.AppError
}
