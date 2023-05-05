package auth

import (
	"thesis/app/services"
	"thesis/config"
)

func Login(email string, password string) bool {
	isLogin := false
	user := services.UserService{}.GetUserByEmail(email)
	if user != nil {
		isLogin = config.CheckPasswordHash(password, user.Password)
	}
	return isLogin
}
