package domain

type Prices struct {
	SKUID       string `db:"sku_id"`
	OutletID    int64  `db:"merchant_id"`
	ProductName string `db:"product_name"`
	Image       string `db:"image"`
	Quantity    int64  `db:"quantity"`
}
