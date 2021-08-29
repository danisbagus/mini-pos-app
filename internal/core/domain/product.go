package domain

type Product struct {
	SKUID       string `db:"sku_id"`
	MerchantID  int64  `db:"merchant_id"`
	ProductName string `db:"product_name"`
	Image       string `db:"image"`
	Quantity    int64  `db:"quantity"`
}

type ProductPrice struct {
	Product
	Price int64 `db:"price"`
}

type ProductOutlet struct {
	OutletID int64 `db:"outlet_id"`
	Product
	Price int64 `db:"price"`
}
