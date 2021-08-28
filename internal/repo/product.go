package repo

import (
	"fmt"
	"strings"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type ProductRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) port.IProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (r ProductRepo) Create(data *domain.ProductPrice, outlets []domain.Outlet) *errs.AppError {

	tx, err := r.db.Begin()
	if err != nil {
		logger.Error("Error when starting new transaction for create product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	insertSql := "insert into products (sku_id, merchant_id, product_name, image, quantity) values (?,?,?,?,?)"

	_, err = r.db.Exec(insertSql, data.SKUID, data.MerchantID, data.ProductName, data.Image, data.Quantity)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while creating new product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	// bulk insert to prices table
	valueStrings := []string{}
	valueArgs := []interface{}{}
	for _, w := range outlets {
		valueStrings = append(valueStrings, "(?, ?, ?)")

		valueArgs = append(valueArgs, data.SKUID)
		valueArgs = append(valueArgs, w.OutletID)
		valueArgs = append(valueArgs, data.Price)
	}

	smt := `insert into prices(sku_id, outlet_id, price) values %s`

	smt = fmt.Sprintf(smt, strings.Join(valueStrings, ","))

	_, err = tx.Exec(smt, valueArgs...)
	if err != nil {
		tx.Rollback()
		logger.Error("Error while creating prices " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	// commit the transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction for product: " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}
