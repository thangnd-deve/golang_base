package database

import (
	"log"
	"thesis/app/entities/users"
)

func Migration() {
	err := Connect().AutoMigrate(users.UserEntity{})
	if err != nil {
		log.Fatal(err.Error())
	}
}
