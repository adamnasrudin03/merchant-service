package service

import (
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"github.com/adamnasrudin03/merchant-service/app/repository"
)

type OutletService interface {
	GetOutletByID(ID int64) (outlet entity.OutletRes, err error)
}

type outletService struct {
	OutletRepository repository.OutletRepository
}

func NewOutletService(OutletRepo repository.OutletRepository) OutletService {
	return &outletService{
		OutletRepository: OutletRepo,
	}
}

func (service *outletService) GetOutletByID(ID int64) (outlet entity.OutletRes, err error) {
	outlet, err = service.OutletRepository.GetByID(ID)
	if err != nil {
		return
	}

	return
}
