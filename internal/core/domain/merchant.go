package domain

type Merchant struct {
	MerchantID        int64  `db:"merchant_id"`
	UserID            int64  `db:"user_id"`
	MerchantName      string `db:"merchant_name"`
	HearOfficeAddress string `db:"head_office_address"`
}
