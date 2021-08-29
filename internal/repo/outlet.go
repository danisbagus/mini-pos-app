package repo

import (
	"database/sql"

	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type OutletRepo struct {
	db *sqlx.DB
}

func NewOutletRepo(db *sqlx.DB) port.IOutletRepo {
	return &OutletRepo{
		db: db,
	}
}

func (r OutletRepo) FindAllByMerchantID(merchantID int64) ([]domain.Outlet, *errs.AppError) {
	outlets := make([]domain.Outlet, 0)

	findAllByMerchantIDSql := "select outlet_id, merchant_id, outlet_name, address from outlets where merchant_id=?"
	err := r.db.Select(&outlets, findAllByMerchantIDSql, merchantID)

	if err != nil {
		logger.Error("Error while quering find all outlet by merchant id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return outlets, nil
}

func (r OutletRepo) FindAll() ([]domain.Outlet, *errs.AppError) {
	outlets := make([]domain.Outlet, 0)

	findAllSql := "select outlet_id, merchant_id, outlet_name, address from outlets"
	err := r.db.Select(&outlets, findAllSql)

	if err != nil {
		logger.Error("Error while quering find all outlet" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return outlets, nil
}

func (r OutletRepo) FindOneByID(outletID int64) (*domain.Outlet, *errs.AppError) {
	var data domain.Outlet

	FindOneSql := "select outlet_id, merchant_id, outlet_name, address from outlets where outlet_id=?"

	err := r.db.Get(&data, FindOneSql, outletID)

	if err != nil {
		logger.Error("Error while get find one by outlet id " + err.Error())
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Outlet not found")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return &data, nil
}

func (r OutletRepo) Create(data *domain.Outlet) (*domain.Outlet, *errs.AppError) {
	insertSql := "insert into outlets (merchant_id, outlet_name, address) values (?,?,?)"

	result, err := r.db.Exec(insertSql, data.MerchantID, data.OutletName, data.Address)
	if err != nil {
		logger.Error("Error while creating new outlet " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while get last insert id for new outlet" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	data.OutletID = id

	return data, nil
}

func (r OutletRepo) Update(outletID int64, data *domain.Outlet) *errs.AppError {

	updateSql := "update outlets set outlet_name=?, address=? where outlet_id=?"

	stmt, err := r.db.Prepare(updateSql)
	if err != nil {
		logger.Error("Error while update outlet " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(data.OutletName, data.Address, outletID)
	if err != nil {
		logger.Error("Error while update outlet " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	return nil
}

func (r OutletRepo) Delete(outletID int64) *errs.AppError {

	deleteSql := "delete from outlets where outlet_id = ?"

	stmt, err := r.db.Prepare(deleteSql)
	if err != nil {
		logger.Error("Error while delete outlet " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}

	_, err = stmt.Exec(outletID)
	if err != nil {
		logger.Error("Error while delete outlet " + err.Error())
		return errs.NewUnexpectedError("Unexpected database error")
	}
	return nil
}
