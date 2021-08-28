package domain

type Outlet struct {
	OutletID   int64  `db:"outlet_id"`
	MerchantID int64  `db:"merchant_id"`
	OutletName string `db:"outlet_name"`
	Address    string `db:"address"`
}
