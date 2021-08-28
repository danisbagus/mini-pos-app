package domain

type Product struct {
	SKUID       string `db:"sku_id"`
	MerchantID  int64  `db:"merchant_id"`
	ProductName string `db:"product_name"`
	Image       string `db:"image"`
	Quantity    int64  `db:"quantity"`
}
