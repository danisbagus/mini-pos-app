package service

import (
	"fmt"
	"time"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type SaleTransactionService struct {
	repo         port.ISaleTransactionRepo
	productRepo  port.IProductRepo
	merchantRepo port.IMerchantRepo
	customerRepo port.ICustomerRepo
	priceRepo    port.IPriceRepo
	outletRepo   port.IOutletRepo
}

func NewSaleTransactionService(repo port.ISaleTransactionRepo, productRepo port.IProductRepo, merchantRepo port.IMerchantRepo, customerRepo port.ICustomerRepo, priceRepo port.IPriceRepo, outletRepo port.IOutletRepo) port.ISaleTransactionService {
	return &SaleTransactionService{
		repo:         repo,
		productRepo:  productRepo,
		merchantRepo: merchantRepo,
		customerRepo: customerRepo,
		priceRepo:    priceRepo,
		outletRepo:   outletRepo,
	}
}

func (r SaleTransactionService) NewTransaction(req *dto.NewSaleTransactionRequest) (*dto.NewSaleTransactionResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// get customer data
	customer, err := r.customerRepo.FindOneByUserID(req.UserID)
	if err != nil {
		return nil, err
	}

	product, err := r.productRepo.FindOne(req.SKUID)
	if err != nil {
		return nil, err
	}

	if product.Quantity < req.Quantity {
		return nil, errs.NewValidationError("Insufficient product quantity")
	}

	if _, err = r.outletRepo.FindOneByID(req.OutletID); err != nil {
		return nil, err

	}

	transactionID := fmt.Sprintf("TS%v", String(6))

	// get price
	priceData, err := r.priceRepo.FindAllBySKUIDAndOutletID(req.SKUID, uint64(req.OutletID))
	if err != nil {
		return nil, err
	}
	totalPrice := priceData.Price * req.Quantity

	if _, err = r.outletRepo.FindOneByID(req.OutletID); err != nil {
		return nil, err
	}

	form := domain.SaleTransaction{
		TransactionID: transactionID,
		CustomerID:    customer.CustomerID,
		SKUID:         req.SKUID,
		OutletID:      req.OutletID,
		Quantity:      req.Quantity,
		TotalPrice:    totalPrice,
		CreatedAt:     time.Now().Format(dbTSLayout),
	}

	err = r.repo.Create(&form)
	if err != nil {
		return nil, err
	}
	response := dto.NewNewSaleTransactionResponse(&form)

	return response, nil
}

func (r SaleTransactionService) GetTransactionReport(userID int64) (*dto.SaleTransactionList, *errs.AppError) {
	merchant, err := r.merchantRepo.FindOneByUserID(userID)
	if err != nil {
		return nil, err
	}

	dataList, err := r.repo.FetchAllByMerchantID(merchant.MerchantID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetSaleTransactionReport(dataList)

	return response, nil
}

func (r SaleTransactionService) GetTransactionReportByProduct(SKUID string, userID int64) (*dto.SaleTransactionList, *errs.AppError) {
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

	response := dto.NewGetSaleTransactionReport(dataList)

	return response, nil
}
