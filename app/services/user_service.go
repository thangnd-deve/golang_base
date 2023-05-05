package services

import (
	"thesis/app/database"
	"thesis/app/entities/users"
	"thesis/app/helpers"
	"thesis/config"
)

func CreateUser(user users.UserDao) (*users.UserDTO, *helpers.ResponseAPIError) {

	userEntity := users.UserEntity{
		Email:    user.Email,
		Password: config.HashPassword(user.Password),
		Name:     user.Name,
	}
	user.Password = userEntity.Password
	userDTO := users.UserDTO{
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
	database.Connect().Create(&userEntity)
	return &userDTO, nil
}

func GetUserByEmail(email string) *users.UserEntity {
	var user *users.UserEntity
	database.Connect().Find(&user, "email = ?", email)
	return user
}
