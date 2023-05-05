package database

import (
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	"log"
	"thesis/app/helpers"
)

func Connect() *gorm.DB {
	dbHost := helpers.Env("DB_HOST", "127.0.0.1")
	dbUserName := helpers.Env("DB_USERNAME", "root")
	dbPassword := helpers.Env("DB_PASSWORD", "root")
	dbDatabase := helpers.Env("DB_DATABASE", "golang")
	dbPort := helpers.Env("DB_PORT", "3306")
	dbDriver := helpers.Env("DB_DRIVER", "mysql")

	dbURL := dbDriver + "://" + dbUserName + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbDatabase

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	return db
}
