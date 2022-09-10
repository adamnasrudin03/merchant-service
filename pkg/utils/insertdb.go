package utils

import (
	"time"

	"github.com/adamnasrudin03/merchant-service/app/entity"
	"gorm.io/gorm"
)

func SeedUsers() *[]entity.User {
	passAdmin1 := HashPassword("admin1")
	passAdmin2 := HashPassword("admin2")

	Users := []entity.User{
		{
			ID:        1,
			Name:      "Admin 1",
			UserName:  "admin1",
			Password:  passAdmin1,
			CreatedAt: time.Now(),
			CreatedBy: 1,
			UpdatedAt: time.Now(),
			UpdatedBy: 1,
		},
		{
			ID:        2,
			Name:      "Admin 2",
			UserName:  "admin2",
			Password:  passAdmin2,
			CreatedAt: time.Now(),
			CreatedBy: 2,
			UpdatedAt: time.Now(),
			UpdatedBy: 2,
		},
	}

	return &Users
}

func SeedMerchants() *[]entity.Merchant {

	Merchants := []entity.Merchant{
		{
			ID:           1,
			UserID:       1,
			MerchantName: "Merchant 1",
			CreatedAt:    time.Now(),
			CreatedBy:    1,
			UpdatedAt:    time.Now(),
			UpdatedBy:    1,
		},
		{
			ID:           2,
			UserID:       2,
			MerchantName: "Merchant 2",
			CreatedAt:    time.Now(),
			CreatedBy:    2,
			UpdatedAt:    time.Now(),
			UpdatedBy:    2,
		},
	}

	return &Merchants
}

func SeedOutlets() *[]entity.Outlet {

	Outlets := []entity.Outlet{
		{
			ID:         1,
			MerchantID: 1,
			OutletName: "Outlet 1",
			CreatedAt:  time.Now(),
			CreatedBy:  1,
			UpdatedAt:  time.Now(),
			UpdatedBy:  1,
		},
		{
			ID:         2,
			MerchantID: 2,
			OutletName: "Outlet 2",
			CreatedAt:  time.Now(),
			CreatedBy:  2,
			UpdatedAt:  time.Now(),
			UpdatedBy:  2,
		},
		{
			ID:         3,
			MerchantID: 1,
			OutletName: "Outlet 3",
			CreatedAt:  time.Now(),
			CreatedBy:  1,
			UpdatedAt:  time.Now(),
			UpdatedBy:  1,
		},
	}

	return &Outlets
}

func InsertDB(db *gorm.DB) {
	tx := db.Begin()
	var Users *[]entity.User
	var Merchants *[]entity.Merchant
	var Outlets *[]entity.Outlet
	var Transactions *[]entity.Transaction

	db.Find(&Users)
	if len(*Users) <= 0 {
		Users = SeedUsers()
		tx.Create(Users)
	}

	db.Find(&Merchants)
	if len(*Merchants) <= 0 {
		Merchants = SeedMerchants()
		tx.Create(Merchants)
	}

	db.Find(&Outlets)
	if len(*Outlets) <= 0 {
		Outlets = SeedOutlets()
		tx.Create(Outlets)
	}

	db.Find(&Transactions)
	if len(*Transactions) <= 0 {
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (1, 1, 1, 2000, '2021-11-01 12:30:04', 1, '2021-11-01 12:30:04',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (2, 1, 1, 2500, '2021-11-01 17:20:14', 1, '2021-11-01 17:20:14',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (3, 1, 1, 4000, '2021-11-02 12:30:04', 1, '2021-11-02 12:30:04',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (4, 1, 1, 1000, '2021-11-04 12:30:04', 1, '2021-11-04 12:30:04',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (5, 1, 1, 7000, '2021-11-05 16:59:30', 1, '2021-11-05 16:59:30',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (6, 1, 3, 2000, '2021-11-02 18:30:04', 1, '2021-11-02 18:30:04',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (7, 1, 3, 2500, '2021-11-03 17:20:14', 1, '2021-11-03 17:20:14',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (8, 1, 3, 4000, '2021-11-04 12:30:04', 1, '2021-11-04 12:30:04',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (9, 1, 3, 1000, '2021-11-04 12:31:04', 1, '2021-11-04 12:31:04',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (10, 1, 3, 7000, '2021-11-05 16:59:30', 1, '2021-11-05 16:59:30',1)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (11, 2, 2, 2000, '2021-11-01 18:30:04', 2, '2021-11-01 18:30:04',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (12, 2, 2, 2500, '2021-11-02 17:20:14', 2, '2021-11-02 17:20:14',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (13, 2, 2, 4000, '2021-11-03 12:30:04', 2, '2021-11-03 12:30:04',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (14, 2, 2, 1000, '2021-11-04 12:31:04', 2, '2021-11-04 12:31:04',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (15, 2, 2, 7000, '2021-11-05 16:59:30', 2, '2021-11-05 16:59:30',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (16, 2, 2, 2000, '2021-11-05 18:30:04', 2, '2021-11-05 18:30:04',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (17, 2, 2, 2500, '2021-11-06 17:20:14', 2, '2021-11-06 17:20:14',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (18, 2, 2, 4000, '2021-11-07 12:30:04', 2, '2021-11-07 12:30:04',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (19, 2, 2, 1000, '2021-11-08 12:31:04', 2, '2021-11-08 12:31:04',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (20, 2, 2, 7000, '2021-11-09 16:59:30', 2, '2021-11-09 16:59:30',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (21, 2, 2, 1000, '2021-11-10 12:31:04', 2, '2021-11-10 12:31:04',2)").Rows()
		_, _ = tx.Raw("insert into transactions (id, merchant_id, outlet_id, bill_total, created_at, created_by, updated_at, updated_by) values (22, 2, 2, 7000, '2021-11-11 16:59:30', 2, '2021-11-11 16:59:30',2)").Rows()
	}
	tx.Commit()
}
