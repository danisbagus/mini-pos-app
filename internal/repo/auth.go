package repo

import (
	"database/sql"
	"time"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/dgrijalva/jwt-go"

	"github.com/jmoiron/sqlx"
)

const ACCESS_TOKEN_DURATION = time.Hour

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) port.IAuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r AuthRepo) FindOne(username string) (*domain.User, *errs.AppError) {
	var login domain.User
	sqlVerify := `select user_id, role, username, password, created_at from users where username = ? `

	err := r.db.Get(&login, sqlVerify, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewAuthenticationError("invalid credentials")
		} else {
			logger.Error("Error while verifying login request from database: " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &login, nil
}

func (r AuthRepo) Verify(token string) *errs.AppError {
	jwtToken, err := jwtTokenFromString(token)
	if err != nil {
		return errs.NewAuthorizationError(err.Error())
	}

	if !jwtToken.Valid {
		return errs.NewAuthorizationError("Invalid token")
	}
	return nil
}

func (r AuthRepo) CreateUserMerchant(data *domain.UserMerchant) (*domain.UserMerchant, *errs.AppError) {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting create new user merchant " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	resultUser, err := tx.Exec(`insert into users (role, username, password, created_at) 
		values (?, ?, ?, ?)`, data.Role, data.Username, data.Password, data.CreatedAt)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while create new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	userID, err := resultUser.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last user id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	resultCustomer, err := tx.Exec(`insert into merchants (user_id, merchant_name, head_office_address) 
	values (?, ?, ?)`, userID, data.MerchantName, data.HearOfficeAddress)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while create new merchant: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	merchantID, err := resultCustomer.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last merchant id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.MerchantID = merchantID
	data.UserID = userID

	return data, nil
}

func (r AuthRepo) CreateUserCustomer(data *domain.UserCustomer) (*domain.UserCustomer, *errs.AppError) {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting create new user customer " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	resultUser, err := tx.Exec(`insert into users (role, username, password, created_at) 
		values (?, ?, ?, ?)`, data.Role, data.Username, data.Password, data.CreatedAt)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while create new user: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	userID, err := resultUser.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last user id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	resultCustomer, err := tx.Exec(`insert into customers (user_id, customer_name, phone) 
	values (?, ?, ?)`, userID, data.CustomerName, data.Phone)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while create new customer: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customerID, err := resultCustomer.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last customer id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.CustomerID = customerID
	data.UserID = userID

	return data, nil
}

func (r AuthRepo) Delete(userID int64) *errs.AppError {

	deleteSql := "delete from users where user_id = ?"

	stmt, err := r.db.Prepare(deleteSql)
	if err != nil {
		logger.Error("Error while delete user " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(userID)
	if err != nil {
		logger.Error("Error while delete user " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}
	return nil
}

func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &domain.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(domain.HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		logger.Error("Error while parsing token: " + err.Error())
		return nil, err
	}
	return token, nil
}
