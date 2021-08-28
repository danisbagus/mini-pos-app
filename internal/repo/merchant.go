package repo

import (
	"database/sql"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type MerchantRepo struct {
	db *sqlx.DB
}

func NewMerchantRepo(db *sqlx.DB) port.IMerchantRepo {
	return &MerchantRepo{
		db: db,
	}
}

func (r MerchantRepo) FindOneByID(merchantID int64) (*domain.UserMerchant, *errs.AppError) {
	var data domain.UserMerchant

	findOneByUserIDSql := `
	select u.*, m.merchant_id, m.merchant_name, m.head_office_address from 
	merchants m inner join users u on u.user_id = m.user_id where m.merchant_id = ?`

	err := r.db.Get(&data, findOneByUserIDSql, merchantID)

	if err != nil {
		logger.Error("Error while get find one mechant by id " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("merchant not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &data, nil
}

func (r MerchantRepo) FindOneByUserID(userID int64) (*domain.UserMerchant, *errs.AppError) {
	var data domain.UserMerchant

	findOneByUserIDSql := `
	select u.*, m.merchant_id, m.merchant_name, m.head_office_address from 
	merchants m inner join users u on u.user_id = m.user_id where m.user_id = ?`

	err := r.db.Get(&data, findOneByUserIDSql, userID)

	if err != nil {
		logger.Error("Error while get find one mechant by user id " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("merchant not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &data, nil
}

func (r MerchantRepo) Update(merchatID int64, data *domain.Merchant) *errs.AppError {
	updateSql := "update merchants set merchant_name=?, head_office_address=? where merchant_id=?"

	stmt, err := r.db.Prepare(updateSql)
	if err != nil {
		logger.Error("Error while update product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(data.MerchantName, data.HearOfficeAddress, merchatID)
	if err != nil {
		logger.Error("Error while update product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}
