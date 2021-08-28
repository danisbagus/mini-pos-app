package repo

import (
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

func (r ProductRepo) Create(data *domain.Product) *errs.AppError {
	insertSql := "insert into products (sku_id, merchant_id, product_name, image, quantity) values (?,?,?,?,?)"

	_, err := r.db.Exec(insertSql, data.SKUID, data.MerchantID, data.ProductName, data.Image, data.Quantity)
	if err != nil {
		logger.Error("Error while creating new product " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}
