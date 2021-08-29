package domain

type Supplier struct {
	SupplierID   int64  `db:"supplier_id"`
	SupplierName string `db:"supplier_name"`
}
