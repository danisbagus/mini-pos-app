package service

import (
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type MerchantService struct {
	repo port.IMerchantRepo
}

func NewMerchantService(repo port.IMerchantRepo) port.IMerchantService {
	return &MerchantService{
		repo: repo,
	}
}

func (r MerchantService) GetDetaiByUserID(userID int64) (*dto.UserMerchantResponse, *errs.AppError) {
	data, err := r.repo.FindOneByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailUserMerchantResponse(data)

	return response, nil
}
