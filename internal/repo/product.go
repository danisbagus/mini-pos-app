package repo

import (
	"database/sql"
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

func (r ProductRepo) FindAllByMerchantID(MerchantID int64) ([]domain.Product, *errs.AppError) {
	products := make([]domain.Product, 0)

	findAllByMerchantIDSql := "select sku_id, merchant_id, product_name, image, quantity from products where merchant_id=?"
	err := r.db.Select(&products, findAllByMerchantIDSql, MerchantID)

	if err != nil {
		logger.Error("Error while quering find all product by merchant id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return products, nil
}

func (r ProductRepo) FindOne(SKUID string) (*domain.Product, *errs.AppError) {
	var data domain.Product

	FindOneSql := "select sku_id, merchant_id, product_name, image, quantity from products where sku_id = ?"

	err := r.db.Get(&data, FindOneSql, SKUID)

	if err != nil {
		logger.Error("Error while get find one by skuid product " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Product not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &data, nil
}

func (r ProductRepo) Update(SKUID string, data *domain.Product) *errs.AppError {
	updateSql := "update products set product_name=?, image=?, quantity=? where sku_id=?"

	stmt, err := r.db.Prepare(updateSql)
	if err != nil {
		logger.Error("Error while update product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(data.ProductName, data.Image, data.Quantity, SKUID)
	if err != nil {
		logger.Error("Error while update product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

func (r ProductRepo) UpdatePrice(SKUID string, outliteID int64, price int64) *errs.AppError {
	updateSql := "update prices set price=? where sku_id=? and outlet_id=?"

	stmt, err := r.db.Prepare(updateSql)
	if err != nil {
		logger.Error("Error while update product price " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(price, SKUID, outliteID)
	if err != nil {
		logger.Error("Error while update product proce" + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil

}

func (r ProductRepo) Delete(SKUID string) *errs.AppError {

	deleteSql := "delete from products where sku_id = ?"

	stmt, err := r.db.Prepare(deleteSql)
	if err != nil {
		logger.Error("Error while delete product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(SKUID)
	if err != nil {
		logger.Error("Error while delete product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}
	return nil
}
