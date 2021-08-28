package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IProductRepo interface {
	Create(data *domain.ProductPrice, outlets []domain.Outlet) *errs.AppError
	FindAllByMerchantID(merchantID int64) ([]domain.Product, *errs.AppError)
	FindOne(SKUID string) (*domain.Product, *errs.AppError)
	Update(SKUID string, data *domain.Product) *errs.AppError
	UpdatePrice(SKUID string, outliteID int64, price int64) *errs.AppError
	Delete(SKUID string) *errs.AppError
}

type IProducService interface {
	NewProduct(data *dto.NewProductRequest) (*dto.NewProductResponse, *errs.AppError)
	GetAllByUserID(userID int64) (*dto.ProductListResponse, *errs.AppError)
	GetDetail(SKUID string) (*dto.ProductResponse, *errs.AppError)
	UpdateProduct(SKUID string, data *dto.UpdateProductRequest) (*dto.UpdateProductResponse, *errs.AppError)
	UpdateProductPrice(SKUID string, data *dto.UpdateProductPriceRequest) *errs.AppError
	RemoveProduct(SKUID string, userID int64) *errs.AppError
}
