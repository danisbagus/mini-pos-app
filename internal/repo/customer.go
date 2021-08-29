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

func (r CustomerRepo) FindAll() ([]domain.Customer, *errs.AppError) {
	customers := make([]domain.Customer, 0)

	findAllSql := "select * from customers"
	err := r.db.Select(&customers, findAllSql)

	if err != nil {
		logger.Error("Error while quering find all customers " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (r CustomerRepo) FindOne(customerID int64) (*domain.Customer, *errs.AppError) {
	var data domain.Customer

	findOneByIDSql := `
	select c.* from 
	customers c where c.customer_id = ?`

	err := r.db.Get(&data, findOneByIDSql, customerID)

	if err != nil {
		logger.Error("Error while get find one mechant by customer id " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &data, nil
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

func (r CustomerRepo) Update(customerID int64, data *domain.Customer) *errs.AppError {
	updateSql := "update customers set customer_name=?, phone=? where customer_id=?"

	stmt, err := r.db.Prepare(updateSql)
	if err != nil {
		logger.Error("Error while update product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(data.CustomerName, data.Phone, customerID)
	if err != nil {
		logger.Error("Error while update product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

func (r CustomerRepo) Delete(customerID int64) *errs.AppError {

	deleteSql := "delete from customers where customer_id = ?"

	stmt, err := r.db.Prepare(deleteSql)
	if err != nil {
		logger.Error("Error while delete product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(customerID)
	if err != nil {
		logger.Error("Error while delete product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}
	return nil
}
