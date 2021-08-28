package service

import (
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

func (r OutletService) GetAllByMerchantID(merchantID int64) (*dto.OutletListResponse, *errs.AppError) {
	// check merchant
	dataList, err := r.repo.FindAllByMerchantID(merchantID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListOutletResponse(dataList)

	return response, nil
}
