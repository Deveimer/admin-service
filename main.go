package main

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/varun-singhh/gofy/pkg/goofy"
)

func main() {

	app := goofy.New()
	_ = connect(app)

	app.Start()
}

func connect(app *goofy.Goofy) *gorm.DB {

	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	pass := os.Getenv("DB_PASSWORD")
	root := os.Getenv("DB_ROOT")
	port := os.Getenv("DAB_PORT")

	dsn := "host=" + host + " user=" + root + " password=" + pass + " dbname=" + name + " port=" + port

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		app.Logger.Errorf("Error while connecting to DB, Error is %v", err)
		return nil
	}
	app.Logger.Log("DB Connected Successfully")

	return db
}
