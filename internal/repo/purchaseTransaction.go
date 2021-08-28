package repo

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type PurchaseTransactionRepo struct {
	db *sqlx.DB
}

func NewPurchaseTransactionRepo(db *sqlx.DB) port.IPurchaseTransactionRepo {
	return &PurchaseTransactionRepo{
		db: db,
	}
}

func (r PurchaseTransactionRepo) Create(data *domain.PurchaseTransaction) *errs.AppError {
	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting new transaction" + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = tx.Exec(`insert into purchase_transactions (transaction_id, merchant_id, sku_id, supplier_id, quantity, total_price, created_at) 
							values (?, ?, ?, ?, ?, ?, ?)`, data.TransactionID, data.MerchantID, data.SKUID, data.SuppierID, data.Quantity, data.TotalPrice, data.CreatedAt)

	if err != nil {
		tx.Rollback()
		logger.Error("Error while create new transaction: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = tx.Exec(`update products set quantity = quantity + ? where sku_id = ?`, data.Quantity, data.SKUID)

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

func (r PurchaseTransactionRepo) FetchAllByMerchantID(merchantID int64) ([]domain.PurchaseTransaction, *errs.AppError) {
	transactions := make([]domain.PurchaseTransaction, 0)

	findAllByMerchantIDSql := `
	select 
		pt.* 
	from purchase_transactions pt 
	where pt.merchant_id=?
	`
	err := r.db.Select(&transactions, findAllByMerchantIDSql, merchantID)

	if err != nil {
		logger.Error("Error while quering find all purchase transaction by merchant id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return transactions, nil
}

func (r PurchaseTransactionRepo) FetchAllBySKUID(SKUID string) ([]domain.PurchaseTransaction, *errs.AppError) {
	transactions := make([]domain.PurchaseTransaction, 0)

	findAllByMerchantIDSql := `
	select 
		pt.* 
	from purchase_transactions pt 
	where pt.sku_id=?
	`
	err := r.db.Select(&transactions, findAllByMerchantIDSql, SKUID)

	if err != nil {
		logger.Error("Error while quering find all purchase transaction by sku id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return transactions, nil
}
