package main

import (
	"log"

	"github.com/marcelofranco/api-tests-examples/contract-test/provider/data"
	"gorm.io/gorm"
)

type Config struct {
	Repo   data.Repository
	Client Client
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

	app := Config{
		Client: NewClient("http://provider"),
	}
	app.setUpRepo(d)

	mux := app.routes()
	mux.Run(port)
}

func (app *Config) setUpRepo(conn *gorm.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}
