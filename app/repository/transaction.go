package repository

import (
	"context"
	"time"

	"github.com/adamnasrudin03/merchant-service/app/dto"
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"gorm.io/gorm"
)

//TransactionRepository is contract what TransactionRepository can do to db
type TransactionRepository interface {
	GetIncomeReport(queryparam dto.ParamTransaction) (result []dto.TransactionRes, total int64, err error)
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

func (repo *TransactionRepo) GetIncomeReport(queryparam dto.ParamTransaction) (result []dto.TransactionRes, total int64, err error) {
	var (
		ctx      context.Context
		dataDB   []entity.TransactionRes
		dataTemp []dto.TransactionRes
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

	err = query.Offset(offset).Limit(queryparam.Limit).Find(&dataDB).Error
	if err != nil {
		return
	}

	startParam, _ := time.Parse("2006-01-02", queryparam.StartAt)
	endParam, _ := time.Parse("2006-01-02", queryparam.EndAt)
	startParam = utils.Date(int(startParam.Year()), int(startParam.Month()), int(startParam.Day()))
	endParam = utils.Date(int(endParam.Year()), int(endParam.Month()), int(endParam.Day()))

	days, _ := utils.DateBetween(startParam, endParam)
	for _, day := range days {
		temp := dto.TransactionRes{
			TransactionDate: day,
		}

		for _, data := range dataDB {
			if day == data.TransactionDate.Format("2006-01-02") {
				temp.MerchantID = data.MerchantID
				temp.MerchantName = data.MerchantName
				temp.OutletID = data.OutletID
				temp.OutletName = data.OutletName
				temp.OmsetTotal = data.OmsetTotal
			} else {
				temp.MerchantID = data.MerchantID
				temp.MerchantName = data.MerchantName
			}
		}

		dataTemp = append(dataTemp, temp)
	}

	start := (queryparam.Page - 1) * queryparam.Limit
	if start > len(dataTemp) {
		start = len(dataTemp)
	}

	end := start + queryparam.Limit
	if end > len(dataTemp) {
		end = len(dataTemp)
	}

	result = dataTemp[start:end]
	total = int64(len(days))

	return
}
