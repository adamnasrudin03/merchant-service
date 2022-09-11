package repository

import (
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"gorm.io/gorm"
)

type OutletRepository interface {
	GetByID(ID int64) (outlet entity.OutletRes, err error)
}

type outletRepo struct {
	DB *gorm.DB
}

func NewOutletRepository(db *gorm.DB) OutletRepository {
	return &outletRepo{
		DB: db,
	}
}

func (repo *outletRepo) GetByID(ID int64) (outlet entity.OutletRes, err error) {
	err = repo.DB.Select(
		"outlets.id", "outlets.merchant_id", "m.merchant_name",
		"m.user_id", "outlets.outlet_name",
		"outlets.created_at", "outlets.created_by",
		"outlets.updated_at", "outlets.updated_by",
	).Joins("join merchants m on m.id = outlets.merchant_id").Where("outlets.id = ?", ID).Find(&outlet).Error
	if err == nil {
		return
	}

	return
}
