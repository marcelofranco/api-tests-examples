package data

import (
	"errors"
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

	runMigrations(DB)

	return DB, nil
}

func runMigrations(d *gorm.DB) error {
	err := d.AutoMigrate(&Student{})

	if err != nil {
		return err
	}

	if d.Migrator().HasTable(&Student{}) {
		if err = d.First(&Student{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var rooms = []Student{
				{
					Name: "Marcelo",
					CPF:  "123.123.123-12",
					RG:   "12.123.123-1",
				},
				{
					Name: "Ana",
					CPF:  "321.321.321-32",
					RG:   "32.321.321-3",
				},
			}
			for _, r := range rooms {
				if err = d.Create(&r).Error; err != nil {
					return err
				}
			}
		}
	}

	return nil
}
