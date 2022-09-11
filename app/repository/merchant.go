package repository

import (
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"gorm.io/gorm"
)

type MerchantRepository interface {
	GetByID(ID int64) (merchant entity.Merchant, err error)
}

type merchantRepo struct {
	DB *gorm.DB
}

func NewMerchantRepository(db *gorm.DB) MerchantRepository {
	return &merchantRepo{
		DB: db,
	}
}

func (repo *merchantRepo) GetByID(ID int64) (merchant entity.Merchant, err error) {
	err = repo.DB.Where("id = ?", ID).Find(&merchant).Error
	if err == nil {
		return
	}

	return
}
