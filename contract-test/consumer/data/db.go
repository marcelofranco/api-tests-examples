package data

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() (*gorm.DB, error) {
	dns := os.Getenv("DATABASE_DSN")
	// dns := "host=postgres port=5432 user=postgres password=password dbname=postgres sslmode=disable timezone=UTC connect_timeout=5"
	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db.")
		return nil, err
	}

	DB.AutoMigrate(&Student{})
	return DB, nil
}
