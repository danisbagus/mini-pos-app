package service

import (
	"fmt"
	"time"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type PurchaseTransactionService struct {
	repo         port.IPurchaseTransactionRepo
	productRepo  port.IProductRepo
	merchantRepo port.IMerchantRepo
}

func NewPurchaseTransactionService(repo port.IPurchaseTransactionRepo, productRepo port.IProductRepo, merchantRepo port.IMerchantRepo) port.IPurchaseTransactionService {
	return &PurchaseTransactionService{
		repo:         repo,
		productRepo:  productRepo,
		merchantRepo: merchantRepo,
	}
}

func (r PurchaseTransactionService) NewTransaction(req *dto.NewPurchaseTransactionRequest) (*dto.NewPurchaseTransactionResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// get merchant data
	merchant, err := r.merchantRepo.FindOneByUserID(req.UserID)
	if err != nil {
		return nil, err
	}

	_, err = r.productRepo.FindOne(req.SKUID)
	if err != nil {
		return nil, err
	}

	transactionID := fmt.Sprintf("TP%v", String(6))

	// validate supplier ...

	form := domain.PurchaseTransaction{
		TransactionID: transactionID,
		MerchantID:    merchant.MerchantID,
		SKUID:         req.SKUID,
		SuppierID:     req.SuppierID,
		Quantity:      req.Quantity,
		TotalPrice:    req.TotalPrice,
		CreatedAt:     time.Now().Format(dbTSLayout),
	}

	err = r.repo.Create(&form)
	if err != nil {
		return nil, err
	}
	response := dto.NewNewPurchaseTransactionResponse(&form)

	return response, nil
}
