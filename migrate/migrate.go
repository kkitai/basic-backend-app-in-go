package main

import (
	"github.com/kkitai/basic-backend-app-in-go/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=basic_backend_app_in_go port=5432 sslmode=disable TimeZone=Asia/Tokyo"

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Telephone{})

	// Create
	db.Create(&model.Telephone{
		OwnerId: 0,
		Number:  "09011112222",
		ICCId:   111111111111111,
	})
	db.Create(&model.Telephone{
		OwnerId: 1,
		Number:  "090222233333",
		ICCId:   222222222222222,
	})

	// Read
	var telephone model.Telephone
	db.First(&telephone, 1) // find product with id 1

	db.First(&telephone, "number = ?", "09011112222") // find product with code l1212
}
