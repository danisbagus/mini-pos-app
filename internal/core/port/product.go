package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IProductRepo interface {
	Create(data *domain.ProductPrice, outlets []domain.Outlet) *errs.AppError
	FindOne(SKUID string) (*domain.Product, *errs.AppError)
}

type IProducService interface {
	NewProduct(data *dto.NewProductRequest) (*dto.NewProductResponse, *errs.AppError)
	GetDetail(SKUID string) (*dto.ProductResponse, *errs.AppError)
}
