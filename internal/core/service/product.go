package service

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"

	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type ProductService struct {
	repo         port.IProductRepo
	merchantRepo port.IMerchantRepo
	outletRepo   port.IOutletRepo
	priceRepo    port.IPriceRepo
}

func NewProductService(repo port.IProductRepo, merchantRepo port.IMerchantRepo, outletRepo port.IOutletRepo, priceRepo port.IPriceRepo) port.IProducService {
	return &ProductService{
		repo:         repo,
		merchantRepo: merchantRepo,
		outletRepo:   outletRepo,
		priceRepo:    priceRepo,
	}
}

func (r ProductService) NewProduct(req *dto.NewProductRequest) (*dto.NewProductResponse, *errs.AppError) {

	defer req.File.Close()

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	// get merchant data
	merchant, err := r.merchantRepo.FindOneByUserID(req.UserID)
	if err != nil {
		return nil, err
	}

	SKUID := fmt.Sprintf("P%v", String(6))

	filePath := fmt.Sprintf("%v:%v/%v", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"), req.Image)

	f, osErr := os.OpenFile(req.Image, os.O_WRONLY|os.O_CREATE, 0666)
	if osErr != nil {
		return nil, errs.NewUnexpectedError("Error when upload file")
	}

	defer f.Close()
	_, _ = io.Copy(f, req.File)

	form := domain.ProductPrice{
		Product: domain.Product{
			SKUID:       SKUID,
			MerchantID:  merchant.MerchantID,
			ProductName: req.ProductName,
			Image:       filePath,
			Quantity:    req.Quantity},
		Price: req.Price,
	}

	// get outlet list
	outletList, err := r.outletRepo.FindAllByMerchantID(merchant.MerchantID)
	if err != nil {
		return nil, err
	}

	err = r.repo.Create(&form, outletList)
	if err != nil {
		return nil, err
	}
	response := dto.NewNewProductResponse(&form)

	return response, nil
}

func (r ProductService) GetDetail(SKUID string) (*dto.ProductResponse, *errs.AppError) {
	data, err := r.repo.FindOne(SKUID)
	if err != nil {
		return nil, err
	}

	prices, err := r.priceRepo.FindAllBySKUID(SKUID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailProductResponse(data, prices)

	return response, nil
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}
