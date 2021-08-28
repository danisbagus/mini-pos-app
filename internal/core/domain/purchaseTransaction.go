package domain

type PurchaseTransaction struct {
	TransactionID string `db:"transaction_id"`
	MerchantID    int64  `db:"merchant_id"`
	SKUID         string `db:"sku_id"`
	SuppierID     int64  `db:"supplier_id"`
	Quantity      int64  `db:"quantity"`
	TotalPrice    int64  `db:"total_price"`
	CreatedAt     string `db:"created_at"`
}
