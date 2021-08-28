package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IOutletRepo interface {
	FindAllByMerchantID(merchantID int64) ([]domain.Outlet, *errs.AppError)
}

type IOutletService interface {
	GetAllByMerchantID(merchantID int64) (*dto.OutletListResponse, *errs.AppError)
}
