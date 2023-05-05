package services

import (
	"thesis/app/database"
	"thesis/app/entities/users"
	"thesis/app/helpers"
	"thesis/config"
)

type UserService struct {
}

func (userService UserService) CreateUser(user users.UserDao) (*users.UserDTO, *helpers.ResponseAPIError) {
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

func (userService UserService) GetUserByEmail(email string) *users.UserEntity {
	var user *users.UserEntity
	database.Connect().Find(&user, "email = ?", email)
	return user
}

func (userService UserService) GetAllUser() []users.UserEntity {
	var user []users.UserEntity
	database.Connect().Find(&user)
	return user
}
