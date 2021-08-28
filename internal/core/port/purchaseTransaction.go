package port

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type IPurchaseTransactionRepo interface {
	Create(data *domain.PurchaseTransaction) *errs.AppError
}

type IPurchaseTransactionService interface {
	NewTransaction(data *dto.NewPurchaseTransactionRequest) (*dto.NewPurchaseTransactionResponse, *errs.AppError)
}
