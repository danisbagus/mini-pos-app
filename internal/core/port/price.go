package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IPriceRepo interface {
	FindAllBySKUID(SKUID string) ([]domain.Prices, *errs.AppError)
}

// type IPriceService interface {
// 	GetAllBySKUID(SKUID int64) (*dto.PriceListResponse, *errs.AppError)
// }
