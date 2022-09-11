package app

import (
	"github.com/adamnasrudin03/merchant-service/app/repository"
	"github.com/adamnasrudin03/merchant-service/app/service"
	"gorm.io/gorm"
)

func WiringRepository(db *gorm.DB) *repository.Repositories {
	return &repository.Repositories{
		Merchant:    repository.NewMerchantRepository(db),
		Outlet:      repository.NewOutletRepository(db),
		Transaction: repository.NewTransactionRepository(db),
		User:        repository.NewUserRepository(db),
	}
}

func WiringService(repo *repository.Repositories) *service.Services {
	return &service.Services{
		Merchant:    service.NewMerchantService(repo.Merchant),
		Outlet:      service.NewOutletService(repo.Outlet),
		Transaction: service.NewTransactionService(repo.Transaction),
		Auth:        service.NewAuthService(repo.User),
	}
}
