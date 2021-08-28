package service

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
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

func (r CustomerService) GetAll() (*dto.CustomerListResponse, *errs.AppError) {
	dataList, err := r.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := dto.NewGetListCustomerResponse(dataList)

	return response, nil
}

func (r CustomerService) GetOne(customerID int64) (*dto.CustomerResponse, *errs.AppError) {

	data, err := r.repo.FindOne(customerID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailCustomerResponse(data)

	return response, nil
}

func (r CustomerService) GetDetailByUserID(userID int64) (*dto.UserCustomerResponse, *errs.AppError) {
	data, err := r.repo.FindOneByUserID(userID)
	if err != nil {
		return nil, err
	}

	response := dto.NewGetDetailUserCustomerResponse(data)

	return response, nil
}

func (r CustomerService) UpdateCustomerByUserID(userID int64, req *dto.UpdateCustomerRequest) *errs.AppError {

	err := req.Validate()
	if err != nil {
		return err
	}

	customer, err := r.GetDetailByUserID(userID)
	if err != nil {
		return err
	}

	form := domain.Customer{
		CustomerName: req.CustomerName,
		Phone:        req.Phone,
	}

	err = r.repo.Update(customer.CustomerID, &form)
	if err != nil {
		return err
	}

	return nil
}
