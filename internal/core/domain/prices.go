package domain

type Prices struct {
	SKUID    string `db:"sku_id"`
	OutletID int64  `db:"outlet_id"`
	Price    int64  `db:"price"`
}
