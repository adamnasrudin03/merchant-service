package dto

type ParamTransaction struct {
	Page       int    `json:"page" valid:"Required"`
	Limit      int    `json:"limit" valid:"Required"`
	Offset     int    `json:"offset"`
	MerchantID int    `json:"merchant_id"`
	OutletID   int    `json:"outlet_id"`
	StartAt    string `json:"start_at" valid:"Required"`
	EndAt      string `json:"end_at" valid:"Required"`
}

type TransactionRes struct {
	MerchantID      int64   `json:"merchant_id"`
	MerchantName    string  `json:"merchant_name"`
	OutletID        int64   `json:"outlet_id"`
	OutletName      string  `json:"outlet_name"`
	OmsetTotal      float64 `json:"omset_total"`
	TransactionDate string  `json:"transaction_date"`
}
