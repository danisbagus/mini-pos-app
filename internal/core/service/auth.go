package service

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"

	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
