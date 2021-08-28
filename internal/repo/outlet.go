package repo

import (
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
