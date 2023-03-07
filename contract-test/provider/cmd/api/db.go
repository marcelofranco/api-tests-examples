package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DBConnect() {
	dns := os.Getenv("DATABASE_DSN")
	DB, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db.")
	}

	DB.AutoMigrate(&Class{})
	DB.AutoMigrate(&StudentClass{})
}
