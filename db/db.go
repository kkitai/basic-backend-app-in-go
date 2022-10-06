package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DRIVER = "postgres"

type Options struct {
	Hostname string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  bool
	TimeZone string
}

type Option func(*Options)

func NewDBConnection(hostname string, port string, user string, password string, dbname string, setters ...Option) (*gorm.DB, error) {
	// Default Options
	args := &Options{
		Hostname: hostname,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
		SSLMode:  true,
		TimeZone: "Asia/Tokyo",
	}

	for _, setter := range setters {
		setter(args)
	}

	var sslmode string
	if args.SSLMode {
		sslmode = "enable"
	} else {
		sslmode = "disable"
	}

	dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s`,
		args.Hostname,
		args.User,
		args.Password,
		args.DBName,
		args.Port,
		sslmode,
		args.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return &gorm.DB{}, err
	}

	return db, nil
}

func SSLMode(sslmode bool) Option {
	return func(args *Options) {
		args.SSLMode = sslmode
	}
}

func TimeZone(timezone string) Option {
	return func(args *Options) {
		args.TimeZone = timezone
	}
}
