package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const ACCESS_TOKEN_DURATION = time.Hour
const HMAC_SAMPLE_SECRET = "miniposapp-secret"

type AccessTokenClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
