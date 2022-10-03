package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DRIVER = "postgres"

type DB struct {
	User       string
	Password   string
	Hostname   string
	DBName     string
	Connection *gorm.DB
}

// TODO: make sslmodel be optional argument
func NewDB(hostname string, port string, user string, password string, dbname string) (*DB, error) {
	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo`, hostname, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return &DB{}, err
	}

	return &DB{
		User:       user,
		Password:   password,
		Hostname:   hostname,
		Connection: db,
	}, nil
}
