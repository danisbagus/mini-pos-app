package repo

import (
	"database/sql"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type SupplierRepo struct {
	db *sqlx.DB
}

func NewSupplierRepo(db *sqlx.DB) port.ISupplierRepo {
	return &SupplierRepo{
		db: db,
	}
}

func (r SupplierRepo) FindAll() ([]domain.Supplier, *errs.AppError) {
	outlets := make([]domain.Supplier, 0)

	findAll := "select s.* from suppliers s"
	err := r.db.Select(&outlets, findAll)

	if err != nil {
		logger.Error("Error while quering find all supplier" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return outlets, nil
}

func (r SupplierRepo) FindOneByID(SupplierID int64) (*domain.Supplier, *errs.AppError) {
	var data domain.Supplier

	FindOneSql := "select s.* from from where supplier_id=?"

	err := r.db.Get(&data, FindOneSql, SupplierID)

	if err != nil {
		logger.Error("Error while get find one by supplier id " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Supplier not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &data, nil
}
