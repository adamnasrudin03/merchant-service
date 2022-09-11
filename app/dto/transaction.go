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
