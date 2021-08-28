package service

import (
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type CustomerService struct {
	repo port.ICustomerRepo
}

func NewCustomerService(repo port.ICustomerRepo) port.ICustomerService {
	return &CustomerService{
		repo: repo,
	}
}

func (r CustomerService) GetDetailByUserID(userID int64) (*dto.UserCustomerResponse, *errs.AppError) {
	data, err := r.repo.FindOneByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailUserCustomerResponse(data)

	return response, nil
}
