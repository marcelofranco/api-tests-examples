package main

import (
	"log"

	"github.com/marcelofranco/api-tests-examples/contract-test/provider/data"
	"gorm.io/gorm"
)

type Client struct {
	DB *gorm.DB
}

const port = ":80"

func main() {
	d, err := data.DBConnect()
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB, err := d.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	app := Client{
		DB: d,
	}

	mux := app.routes()
	mux.Run(port)
}
