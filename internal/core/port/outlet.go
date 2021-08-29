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
}

type IOutletService interface {
	NewOutlet(data *dto.NewOutletRequest) (*dto.NewOutletResponse, *errs.AppError)
	GetAllByMerchantID(merchantID int64) (*dto.OutletListResponse, *errs.AppError)
}
