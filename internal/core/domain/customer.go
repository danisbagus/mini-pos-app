package domain

type Customer struct {
	CustomerID   int64  `db:"customer_id"`
	UserID       int64  `db:"user_id"`
	CustomerName string `db:"customer_name"`
	Phone        string `db:"phone"`
}
