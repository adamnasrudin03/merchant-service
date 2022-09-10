package entity

import "time"

type Merchant struct {
	ID           int64     `gorm:"primary_key:auto_increment" json:"id"`
	UserID       int64     `gorm:"type:int(40);NOT NULL" json:"user_id"`
	MerchantName string    `gorm:"type:varchar(40);NOT NULL" json:"merchant_name"`
	CreatedAt    time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy    int64     `gorm:"type:bigint(20);NOT NULL" json:"created_by"`
	UpdatedAt    time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy    int64     `gorm:"type:bigint(20);NOT NULL" json:"updated_by"`
}
