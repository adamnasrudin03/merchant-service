package entity

import "time"

type Outlet struct {
	ID         int64     `gorm:"primary_key:auto_increment" json:"id"`
	MerchantID int64     `gorm:"type:bigint(20);NOT NULL" json:"merchant_id"`
	OutletName string    `gorm:"type:varchar(40);NOT NULL" json:"outlet_name"`
	CreatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy  int64     `gorm:"type:bigint(20);NOT NULL" json:"created_by"`
	UpdatedAt  time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy  int64     `gorm:"type:bigint(20);NOT NULL" json:"updated_by"`
}

type OutletRes struct {
	ID           int64     `json:"id"`
	UserID       int64     `json:"user_id"`
	MerchantID   int64     `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	OutletName   string    `json:"outlet_name"`
	CreatedAt    time.Time `json:"created_at"`
	CreatedBy    int64     `json:"created_by"`
	UpdatedAt    time.Time `json:"updated_at"`
	UpdatedBy    int64     `json:"updated_by"`
}

func (OutletRes) TableName() string {
	return "outlets"
}
