package dto

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserMerchantResponse struct {
	UserID            int64  `json:"user_id"`
	Role              string `json:"role"`
	Username          string `json:"username"`
	MerchantID        int64  `json:"merchant_id"`
	MerchantName      string `json:"merchant_name"`
	HearOfficeAddress string `json:"head_office_address"`
}

type UpdateMerchanteRequest struct {
	MerchantName      string `json:"merchant_name"`
	HearOfficeAddress string `json:"head_office_address"`
}

type MerchantResponse struct {
	MerchantID        int64  `json:"merchant_id"`
	UserID            int64  `json:"user_id"`
	MerchantName      string `json:"merchant_name"`
	HearOfficeAddress string `json:"head_office_address"`
}

type MerchantListResponse struct {
	Merchant []MerchantResponse `json:"data"`
}

func NewGetDetailUserMerchantResponse(data *domain.UserMerchant) *UserMerchantResponse {
	result := &UserMerchantResponse{
		UserID:            data.UserID,
		Role:              data.Role,
		Username:          data.Username,
		MerchantID:        data.MerchantID,
		MerchantName:      data.MerchantName,
		HearOfficeAddress: data.HearOfficeAddress,
	}
	return result
}

func NewGetListMerchantResponse(data []domain.Merchant) *MerchantListResponse {
	dataList := make([]MerchantResponse, len(data))

	for k, v := range data {
		dataList[k] = MerchantResponse{
			MerchantID:        v.MerchantID,
			MerchantName:      v.MerchantName,
			UserID:            v.UserID,
			HearOfficeAddress: v.HearOfficeAddress,
		}
	}
	return &MerchantListResponse{Merchant: dataList}
}

func NewGetDetailMerchantResponse(data *domain.UserMerchant) *MerchantResponse {
	result := &MerchantResponse{
		MerchantID:        data.MerchantID,
		UserID:            data.UserID,
		MerchantName:      data.MerchantName,
		HearOfficeAddress: data.HearOfficeAddress,
	}
	return result
}

func (r UpdateMerchanteRequest) Validate() *errs.AppError {
	if err := validation.Validate(r.MerchantName, validation.Required); err != nil {
		return errs.NewBadRequestError("MerchantName is required")
	}

	if err := validation.Validate(r.HearOfficeAddress, validation.Required); err != nil {
		return errs.NewBadRequestError("HearOfficeAddress quantity is required")
	}
	return nil
}
