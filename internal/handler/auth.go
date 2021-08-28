package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/internal/dto"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/dgrijalva/jwt-go"
)

type AuthHandler struct {
	Service port.IAuthService
}

func (rc AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logger.Error("Error while decoding login request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, appErr := rc.Service.Login(loginRequest)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, *token)

}

func (rc AuthHandler) RegisterMerchant(w http.ResponseWriter, r *http.Request) {
	var registerRequest dto.RegisterMerchantRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		logger.Error("Error while decoding register merchant request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, appErr := rc.Service.RegisterMerchant(&registerRequest)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, *token)

}

func (rc AuthHandler) RegisterCustomer(w http.ResponseWriter, r *http.Request) {
	var registerRequest dto.RegisterCustomerRequest
	if err := json.NewDecoder(r.Body).Decode(&registerRequest); err != nil {
		logger.Error("Error while decoding register customer request: " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, appErr := rc.Service.RegisterCustomer(&registerRequest)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, *token)

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func GetClaimData(r *http.Request) (*domain.AccessTokenClaims, *errs.AppError) {
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer")
	var token string

	if len(splitToken) == 2 {
		token = strings.TrimSpace(splitToken[1])
	} else {
		logger.Error("Error while split token")
		return nil, errs.NewAuthorizationError("Invalid token")
	}

	jwtToken, err := jwt.ParseWithClaims(token, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, errs.NewAuthorizationError(err.Error())
	}

	if !jwtToken.Valid {
		return nil, errs.NewAuthorizationError("Invalid token")
	}

	claims := jwtToken.Claims.(*domain.AccessTokenClaims)
	return claims, nil
}
