package gormdb

import (
	"fmt"

	"github.com/adamnasrudin03/merchant-service/app/configs"
	"github.com/adamnasrudin03/merchant-service/app/entity"
	"github.com/adamnasrudin03/merchant-service/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDbConnection is creating a new connection to our database
func SetupDbConnection() *gorm.DB {
	configs := configs.GetInstance()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configs.Dbconfig.Username,
		configs.Dbconfig.Password,
		configs.Dbconfig.Host,
		configs.Dbconfig.Port,
		configs.Dbconfig.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	if configs.Dbconfig.DbIsMigrate {
		//auto migration entity db
		db.AutoMigrate(
			&entity.User{},
			&entity.Merchant{},
			&entity.Outlet{},
			&entity.Transaction{},
		)
		utils.InsertDB(db)
	}

	fmt.Println("Connection Database Success!")
	return db
}

//CloseDbConnection method is closing a connection between your app and your db
func CloseDbConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
