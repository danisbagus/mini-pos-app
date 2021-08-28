package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type ISaleTransactionRepo interface {
	Create(data *domain.SaleTransaction) *errs.AppError
	FetchAllByMerchantID(merchantID int64) ([]domain.SaleTransaction, *errs.AppError)
	FetchAllBySKUID(SKUID string) ([]domain.SaleTransaction, *errs.AppError)
}

type ISaleTransactionService interface {
	NewTransaction(data *dto.NewSaleTransactionRequest) (*dto.NewSaleTransactionResponse, *errs.AppError)
	GetTransactionReport(userID int64) (*dto.SaleTransactionList, *errs.AppError)
	GetTransactionReportByProduct(SKUID string, userID int64) (*dto.SaleTransactionList, *errs.AppError)
}
