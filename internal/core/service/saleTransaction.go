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
	customerRepo port.ICustomerRepo
	priceRepo    port.IPriceRepo
}

func NewSaleTransactionService(repo port.ISaleTransactionRepo, productRepo port.IProductRepo, customerRepo port.ICustomerRepo, priceRepo port.IPriceRepo) port.ISaleTransactionService {
	return &SaleTransactionService{
		repo:         repo,
		productRepo:  productRepo,
		customerRepo: customerRepo,
		priceRepo:    priceRepo,
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

	transactionID := fmt.Sprintf("TS%v", String(6))

	// get price
	priceData, err := r.priceRepo.FindAllBySKUIDAndOutletID(req.SKUID, uint64(req.OutletID))
	if err != nil {
		return nil, err
	}
	totalPrice := priceData.Price * req.Quantity

	// validate outlite ...

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
