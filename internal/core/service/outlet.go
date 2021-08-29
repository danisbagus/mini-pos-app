package service

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type OutletService struct {
	repo         port.IOutletRepo
	merchantRepo port.IMerchantRepo
}

func NewOutletService(repo port.IOutletRepo, merchantRepo port.IMerchantRepo) port.IOutletService {
	return &OutletService{
		repo:         repo,
		merchantRepo: merchantRepo,
	}
}

func (r OutletService) NewOutlet(req *dto.NewOutletRequest) (*dto.NewOutletResponse, *errs.AppError) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	merchant, err := r.merchantRepo.FindOneByUserID(req.UserID)
	if err != nil {
		return nil, err
	}

	form := domain.Outlet{
		OutletName: req.OutletName,
		MerchantID: merchant.MerchantID,
		Address:    req.Address,
	}

	newData, err := r.repo.Create(&form)
	if err != nil {
		return nil, err
	}
	response := dto.NewNewOutletResponse(newData)

	return response, nil
}

func (r OutletService) GetAllByMerchantID(merchantID int64) (*dto.OutletListResponse, *errs.AppError) {
	// check merchant
	dataList, err := r.repo.FindAllByMerchantID(merchantID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListOutletResponse(dataList)

	return response, nil
}

func (r OutletService) GetAll() (*dto.OutletListResponse, *errs.AppError) {
	dataList, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListOutletResponse(dataList)

	return response, nil
}

func (r OutletService) GetAllByUserID(userID int64) (*dto.OutletListResponse, *errs.AppError) {

	merchant, err := r.merchantRepo.FindOneByUserID(userID)
	if err != nil {
		return nil, err
	}
	// check merchant
	dataList, err := r.repo.FindAllByMerchantID(merchant.MerchantID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListOutletResponse(dataList)

	return response, nil
}

func (r OutletService) UpdateOutlet(outletID int64, req *dto.NewOutletRequest) *errs.AppError {

	err := req.Validate()
	if err != nil {
		return err
	}

	outlet, err := r.repo.FindOneByID(outletID)
	if err != nil {
		return err
	}

	merchant, err := r.merchantRepo.FindOneByUserID(req.UserID)
	if err != nil {
		return err
	}

	if outlet.MerchantID != merchant.MerchantID {
		return errs.NewBadRequestError("Cannot update outlet of another merchant")
	}

	form := domain.Outlet{
		OutletName: req.OutletName,
		Address:    req.Address,
	}

	err = r.repo.Update(outletID, &form)
	if err != nil {
		return err
	}

	return nil
}

func (r OutletService) RemoveOutlet(outletID int64, userID int64) *errs.AppError {
	outlet, err := r.repo.FindOneByID(outletID)
	if err != nil {
		return err
	}

	merchant, err := r.merchantRepo.FindOneByUserID(userID)
	if err != nil {
		return err
	}

	if outlet.MerchantID != merchant.MerchantID {
		return errs.NewBadRequestError("Cannot remove outlet of another merchant")
	}

	err = r.repo.Delete(outletID)
	if err != nil {
		return err
	}
	return nil
}
