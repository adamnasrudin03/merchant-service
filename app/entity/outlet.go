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
