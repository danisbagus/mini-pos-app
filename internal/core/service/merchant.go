package service

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type MerchantService struct {
	repo    port.IMerchantRepo
	autRepo port.IAuthRepo
}

func NewMerchantService(repo port.IMerchantRepo, autRepo port.IAuthRepo) port.IMerchantService {
	return &MerchantService{
		repo:    repo,
		autRepo: autRepo,
	}
}

func (r MerchantService) GetAll() (*dto.MerchantListResponse, *errs.AppError) {
	dataList, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListMerchantResponse(dataList)

	return response, nil
}

func (r MerchantService) GetOne(merchantID int64) (*dto.MerchantResponse, *errs.AppError) {

	data, err := r.repo.FindOneByID(merchantID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailMerchantResponse(data)

	return response, nil
}

func (r MerchantService) GetDetailByUserID(userID int64) (*dto.UserMerchantResponse, *errs.AppError) {

	data, err := r.repo.FindOneByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailUserMerchantResponse(data)

	return response, nil
}

func (r MerchantService) UpdateMerchantByID(merchantID int64, req *dto.UpdateMerchanteRequest) *errs.AppError {

	err := req.Validate()
	if err != nil {
		return err
	}

	merchant, err := r.GetOne(merchantID)
	if err != nil {
		return err
	}

	form := domain.Merchant{
		MerchantName:      req.MerchantName,
		HearOfficeAddress: req.HearOfficeAddress,
	}

	err = r.repo.Update(merchant.MerchantID, &form)
	if err != nil {
		return err
	}

	return nil
}

func (r MerchantService) UpdateMerchantByUserID(userID int64, req *dto.UpdateMerchanteRequest) *errs.AppError {

	err := req.Validate()
	if err != nil {
		return err
	}

	merchant, err := r.GetDetailByUserID(userID)
	if err != nil {
		return err
	}

	form := domain.Merchant{
		MerchantName:      req.MerchantName,
		HearOfficeAddress: req.HearOfficeAddress,
	}

	err = r.repo.Update(merchant.MerchantID, &form)
	if err != nil {
		return err
	}

	return nil
}

func (r MerchantService) RemoveMerchant(merchantID int64) *errs.AppError {
	merchant, err := r.GetOne(merchantID)
	if err != nil {
		return err
	}

	err = r.autRepo.Delete(merchant.UserID)
	if err != nil {
		return err
	}
	return nil
}
