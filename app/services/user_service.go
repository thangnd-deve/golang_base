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

func (userService UserService) UpdateUser(userId int64, update users.UserUpdate) *users.UserEntity {
	var user users.UserEntity
	err := database.Connect().Where("id=?", userId).First(&user)
	if err.Error != nil {
		return nil
	}
	user.Name = update.Name
	user.Email = update.Email
	database.Connect().Updates(&user)
	return &user

}

func (userService UserService) DeleteUser(userId int64) bool {
	var user users.UserEntity
	err := database.Connect().Where("id=?", userId).First(&user)
	if err.Error != nil {
		return false
	}
	database.Connect().Delete(&user)
	return true

}
