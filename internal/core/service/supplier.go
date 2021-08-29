package service

import (
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type SupplierService struct {
	repo port.ISupplierRepo
}

func NewSupplierService(repo port.ISupplierRepo) port.ISupplierService {
	return &SupplierService{
		repo: repo,
	}
}

func (r SupplierService) GetAll() (*dto.SupplierListResponse, *errs.AppError) {
	dataList, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListSupplierResponse(dataList)

	return response, nil
}
