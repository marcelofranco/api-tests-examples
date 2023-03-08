package main

import (
	"errors"
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

	runMigrations(DB)
}

func runMigrations(d *gorm.DB) error {
	err := d.AutoMigrate(&Class{}, &StudentClass{})

	if err != nil {
		return err
	}

	if d.Migrator().HasTable(&Class{}) {
		if err = d.First(&Class{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var rooms = []Class{
				{
					Discipline: "Math",
					Day:        "Friday",
					Hour:       "09:00",
				},
				{
					Discipline: "History",
					Day:        "Tuesday",
					Hour:       "10:00",
				},
			}
			for _, r := range rooms {
				if err = d.Create(&r).Error; err != nil {
					return err
				}
			}
		}
	}

	if d.Migrator().HasTable(&StudentClass{}) {
		if err = d.First(&StudentClass{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			var rooms = []StudentClass{
				{
					ClassID:   1,
					StudentId: 1,
				},
				{
					ClassID:   2,
					StudentId: 1,
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
