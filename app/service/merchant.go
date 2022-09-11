package service

import (
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"github.com/adamnasrudin03/merchant-service/app/repository"
)

type MerchantService interface {
	GetMerchantByID(ID int64) (merchant entity.Merchant, err error)
}

type merchantService struct {
	merchantRepository repository.MerchantRepository
}

func NewMerchantService(merchantRepo repository.MerchantRepository) MerchantService {
	return &merchantService{
		merchantRepository: merchantRepo,
	}
}

func (service *merchantService) GetMerchantByID(ID int64) (merchant entity.Merchant, err error) {
	merchant, err = service.merchantRepository.GetByID(ID)
	if err != nil {
		return
	}

	return
}
