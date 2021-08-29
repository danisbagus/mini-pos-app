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
	suppierRepo  port.ISupplierRepo
}

func NewPurchaseTransactionService(repo port.IPurchaseTransactionRepo, productRepo port.IProductRepo, merchantRepo port.IMerchantRepo, suppierRepo port.ISupplierRepo) port.IPurchaseTransactionService {
	return &PurchaseTransactionService{
		repo:         repo,
		productRepo:  productRepo,
		merchantRepo: merchantRepo,
		suppierRepo:  suppierRepo,
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
	if _, err := r.suppierRepo.FindOneByID(req.SuppierID); err != nil {
		return nil, err
	}

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

func (r PurchaseTransactionService) GetTransactionReport(userID int64) (*dto.PurchaseTransactionList, *errs.AppError) {
	merchant, err := r.merchantRepo.FindOneByUserID(userID)
	if err != nil {
		return nil, err
	}

	dataList, err := r.repo.FetchAllByMerchantID(merchant.MerchantID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetPurchaseTransactionReport(dataList)

	return response, nil
}

func (r PurchaseTransactionService) GetTransactionReportByProduct(SKUID string, userID int64) (*dto.PurchaseTransactionList, *errs.AppError) {
	merchant, err := r.merchantRepo.FindOneByUserID(userID)
	if err != nil {
		return nil, err
	}

	product, err := r.productRepo.FindOne(SKUID)
	if err != nil {
		return nil, err
	}

	if product.MerchantID != merchant.MerchantID {
		return nil, errs.NewBadRequestError("Cannot get product data of another merchant")
	}

	dataList, err := r.repo.FetchAllBySKUID(SKUID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetPurchaseTransactionReport(dataList)

	return response, nil
}
