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

type UserMerchant struct {
	User
	MerchantID        int64  `db:"merchant_id"`
	MerchantName      string `db:"merchant_name"`
	HearOfficeAddress string `db:"head_office_address"`
}

type UserCustomer struct {
	User
	CustomerID   int64  `db:"customer_id"`
	CustomerName string `db:"customer_name"`
	Phone        string `db:"phone"`
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
