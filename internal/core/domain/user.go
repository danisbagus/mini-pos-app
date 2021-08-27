package domain

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	UserID    int64  `db:"user_id"`
	Role      string `db:"role"`
	Username  string `db:"username"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
}

func (r User) ClaimsForAccessToken() AccessTokenClaims {
	return AccessTokenClaims{
		UserID: r.UserID,
		Role:   r.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix(),
		},
	}
}
