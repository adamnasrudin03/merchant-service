package entity

import "time"

type Transaction struct {
	ID         int64     `gorm:"primary_key:auto_increment" json:"id"`
	MerchantID int64     `gorm:"type:bigint(20);NOT NULL" json:"merchant_id"`
	OutletID   int64     `gorm:"type:bigint(20);NOT NULL" json:"outlet_id"`
	BillTotal  float64   `gorm:"type:double;NOT NULL" json:"bill_total"`
	CreatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy  int64     `gorm:"type:bigint(20);NOT NULL" json:"created_by"`
	UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy  int64     `gorm:"type:bigint(20);NOT NULL" json:"updated_by"`
}

type TransactionRes struct {
	MerchantID      int64     `json:"merchant_id"`
	MerchantName    string    `json:"merchant_name"`
	OutletID        int64     `json:"outlet_id"`
	OutletName      string    `json:"outlet_name"`
	OmsetTotal      float64   `json:"omset_total"`
	TransactionDate time.Time `json:"transaction_date"`
}

func (TransactionRes) TableName() string {
	return "transactions"
}
