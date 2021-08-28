package repo

import (
	"github.com/danisbagus/mini-pos-app/internal/core/domain"
	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
	"github.com/danisbagus/mini-pos-app/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type PriceRepo struct {
	db *sqlx.DB
}

func NewPriceRepo(db *sqlx.DB) port.IPriceRepo {
	return &PriceRepo{
		db: db,
	}
}

func (r PriceRepo) FindAllBySKUID(SKUID string) ([]domain.Prices, *errs.AppError) {
	outlets := make([]domain.Prices, 0)

	findAllBySKUIDSql := "select sku_id, outlet_id, price from prices where sku_id=?"
	err := r.db.Select(&outlets, findAllBySKUIDSql, SKUID)

	if err != nil {
		logger.Error("Error while quering find all prices by sku id " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return outlets, nil
}
