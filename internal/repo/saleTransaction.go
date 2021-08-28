package repo

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type SaleTransactionRepo struct {
	db *sqlx.DB
}

func NewSaleTransactionRepo(db *sqlx.DB) port.ISaleTransactionRepo {
	return &SaleTransactionRepo{
		db: db,
	}
}

func (r SaleTransactionRepo) Create(data *domain.SaleTransaction) *errs.AppError {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting new transaction" + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = tx.Exec(`insert into sale_transactions (transaction_id, customer_id, sku_id, outlet_id, quantity, total_price, created_at) 
							values (?, ?, ?, ?, ?, ?, ?)`, data.TransactionID, data.CustomerID, data.SKUID, data.OutletID, data.Quantity, data.TotalPrice, data.CreatedAt)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while create new transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = tx.Exec(`update products set quantity = quantity - ? where sku_id = ?`, data.Quantity, data.SKUID)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while update product quantity: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for product: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}
