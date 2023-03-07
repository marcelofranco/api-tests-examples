package data

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConnect() (*gorm.DB, error) {
	dsn := "host=postgres port=5432 user=postgres password=password dbname=postgres sslmode=disable timezone=UTC connect_timeout=5"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db.")
		return nil, err
	}

	DB.AutoMigrate(&Student{})
	return DB, nil
}
