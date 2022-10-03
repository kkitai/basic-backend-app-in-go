package main

import (
	"fmt"
	"os"

	"github.com/kelseyhightower/envconfig"
	"github.com/kkitai/basic-backend-app-in-go/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Env struct {
	DBHost     string `default:"localhost"`
	DBPort     string `default:"5432"`
	DBName     string `required:"true"`
	DBUser     string `required:"true"`
	DBPassword string `required:"true"`
}

func main() {
	var env Env
	if err := envconfig.Process("myapp", &env); err != nil {
		fmt.Fprintf(os.Stderr, "failed to load environment variables: %s\n", err.Error())
		os.Exit(1)
	}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo`, env.DBHost, env.DBUser, env.DBPassword, env.DBName, env.DBPort)

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
