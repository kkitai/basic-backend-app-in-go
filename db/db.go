package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DRIVER = "postgres"

// TODO: make sslmodel and timezone be optional argument
func NewDBConnection(hostname string, port string, user string, password string, dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo`, hostname, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}
