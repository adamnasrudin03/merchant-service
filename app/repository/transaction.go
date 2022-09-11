package repository

import (
	"context"

	"github.com/adamnasrudin03/merchant-service/app/dto"
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"gorm.io/gorm"
)

//TransactionRepository is contract what TransactionRepository can do to db
type TransactionRepository interface {
	GetIncomeReport(queryparam dto.ParamTransaction) (result []entity.TransactionRes, total int64, err error)
}

type TransactionRepo struct {
	DB *gorm.DB
}

//NewTransactionRepository is creates a new instance of TransactionRepository
func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepo{
		DB: db,
	}
}

func (repo *TransactionRepo) GetIncomeReport(queryparam dto.ParamTransaction) (result []entity.TransactionRes, total int64, err error) {
	var (
		ctx context.Context
	)

	offset := queryparam.Limit * (queryparam.Page - 1)

	query := repo.DB.WithContext(ctx).
		Select("transactions.merchant_id", "m.merchant_name",
			"transactions.outlet_id", "o.outlet_name",
			"IFNULL(sum(transactions.bill_total), 0) as omset_total",
			"date(transactions.created_at) as transaction_date",
		).
		Joins("join merchants m on m.id = transactions.merchant_id").
		Joins("join outlets o on o.id = transactions.outlet_id")

	if queryparam.MerchantID != 0 {
		query = query.Where("transactions.merchant_id = ? ", queryparam.MerchantID)
	}

	if queryparam.OutletID != 0 {
		query = query.Where("transactions.outlet_id = ? ", queryparam.OutletID)
	}

	query = query.Where("date(transactions.created_at) BETWEEN ? AND ?", queryparam.StartAt, queryparam.EndAt).
		Group("date(transactions.created_at)").
		Order("transactions.created_at asc")

	err = query.Model(&entity.TransactionRes{}).Count(&total).Error
	if err != nil {
		return
	}

	err = query.Offset(offset).Limit(queryparam.Limit).Find(&result).Error
	if err != nil {
		return
	}

	return
}
