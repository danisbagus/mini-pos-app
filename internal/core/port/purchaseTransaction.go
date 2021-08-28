package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IPurchaseTransactionRepo interface {
	Create(data *domain.PurchaseTransaction) *errs.AppError
	FetchAllByMerchantID(merchantID int64) ([]domain.PurchaseTransaction, *errs.AppError)
	FetchAllBySKUID(SKUID string) ([]domain.PurchaseTransaction, *errs.AppError)
}

type IPurchaseTransactionService interface {
	NewTransaction(data *dto.NewPurchaseTransactionRequest) (*dto.NewPurchaseTransactionResponse, *errs.AppError)
	GetTransactionReport(userID int64) (*dto.PurchaseTransactionList, *errs.AppError)
	GetTransactionReportByProduct(SKUID string, userID int64) (*dto.PurchaseTransactionList, *errs.AppError)
}
