package repository

import (
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	VerifyCredential(email string, password string) interface{}
}

type userRepo struct {
	DB *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		DB: db,
	}
}

func (repo *userRepo) VerifyCredential(username string, password string) interface{} {
	var user entity.User
	res := repo.DB.Where("user_name = ?", username).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}
