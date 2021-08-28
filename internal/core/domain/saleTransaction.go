package domain

type SaleTransaction struct {
	TransactionID string `db:"transaction_id"`
	CustomerID    int64  `db:"customer_id"`
	SKUID         string `db:"sku_id"`
	OutletID      int64  `db:"outlet_id"`
	Quantity      int64  `db:"quantity"`
	TotalPrice    int64  `db:"total_price"`
	CreatedAt     string `db:"created_at"`
}
