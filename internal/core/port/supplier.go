package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type ISupplierRepo interface {
	FindAll() ([]domain.Supplier, *errs.AppError)
	FindOneByID(outletID int64) (*domain.Supplier, *errs.AppError)
}

type ISupplierService interface {
	GetAll() (*dto.SupplierListResponse, *errs.AppError)
}
