package repo

import (
	"database/sql"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type CustomerRepo struct {
	db *sqlx.DB
}

func NewCustomerRepo(db *sqlx.DB) port.ICustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

func (r CustomerRepo) FindOneByUserID(userID int64) (*domain.UserCustomer, *errs.AppError) {
	var data domain.UserCustomer

	findOneByUserIDSql := `
	select u.*, c.customer_id, c.customer_name, c.phone from 
	customers c inner join users u on u.user_id = u.user_id where u.user_id = ?`

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
