package entity

import "time"

type User struct {
	ID        int64     `gorm:"primary_key:auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(45);DEFAULT NULL" json:"name"`
	UserName  string    `gorm:"column:user_name;type:varchar(45);DEFAULT NULL" json:"user_name"`
	Password  string    `gorm:"type:varchar(255);DEFAULT NULL" json:"-"`
	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP" json:"created_at"`
	CreatedBy int64     `gorm:"type:bigint(20);not null" json:"created_by"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`
	UpdatedBy int64     `gorm:"type:bigint(20);not null" json:"updated_by"`
}
