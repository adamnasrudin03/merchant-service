package service

import (
	"math"

	"github.com/adamnasrudin03/merchant-service/app/dto"
	"github.com/adamnasrudin03/merchant-service/app/repository"
)

//TransactionService is a contract about something that this service can do
type TransactionService interface {
	GetIncomeReport(queryparam dto.ParamTransaction) (result dto.ResponseList, err error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

//NewTransactionService creates a new instance of TransactionService
func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepo,
	}
}
func (u *transactionService) GetIncomeReport(queryparam dto.ParamTransaction) (result dto.ResponseList, err error) {
	result.Limit = int64(queryparam.Limit)
	result.Page = int64(queryparam.Page)

	result.Data, result.Total, err = u.transactionRepository.GetIncomeReport(queryparam)
	if err != nil {
		return result, err
	}

	result.LastPage = int64(math.Ceil(float64(result.Total) / float64(queryparam.Limit)))

	return result, nil
}
