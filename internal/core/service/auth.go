package service

import (
	"time"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"

	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AuthService struct {
	repo port.IAuthRepo
}

func NewAuthServie(repo port.IAuthRepo) port.IAuthService {
	return &AuthService{
		repo: repo,
	}
}

func (r AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	var appErr *errs.AppError
	var login *domain.User

	err := req.Validate()

	if err != nil {
		return nil, err
	}

	if login, appErr = r.repo.FindOne(req.Username); appErr != nil {
		return nil, appErr
	}

	match := CheckPasswordHash(req.Password, login.Password)
	if !match {
		return nil, errs.NewAuthenticationError("invalid credentials")
	}

	claims := login.ClaimsForAccessToken()

	authToken := domain.NewAuthToken(claims)

	var accessToken string
	if accessToken, appErr = authToken.NewAccessToken(); appErr != nil {
		return nil, appErr
	}

	return &dto.LoginResponse{AccessToken: accessToken}, nil

}

func (r AuthService) RegisterMerchant(req *dto.RegisterMerchantRequest) (*dto.RegisterMerchantResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	hashPassword, _ := HashPassword(req.Password)

	form := domain.UserMerchant{
		User:              domain.User{Role: "MERCHANT", Username: req.Username, Password: hashPassword, CreatedAt: time.Now().Format(dbTSLayout)},
		MerchantName:      req.MerchantName,
		HearOfficeAddress: req.HearOfficeAddress,
	}

	newData, err := r.repo.CreateUserMerchant(&form)
	if err != nil {
		return nil, err
	}
	response := dto.NewRegisterUserMerchantResponse(newData)

	return response, nil

}

func (r AuthService) RegisterCustomer(req *dto.RegisterCustomerRequest) (*dto.RegisterCustomerResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	hashPassword, _ := HashPassword(req.Password)

	form := domain.UserCustomer{
		User:         domain.User{Role: "CUSTOMER", Username: req.Username, Password: hashPassword, CreatedAt: time.Now().Format(dbTSLayout)},
		CustomerName: req.CustomerName,
		Phone:        req.Phone,
	}

	newData, err := r.repo.CreateUserCustomer(&form)
	if err != nil {
		return nil, err
	}
	response := dto.NewRegisterUserCustomerResponse(newData)

	return response, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
