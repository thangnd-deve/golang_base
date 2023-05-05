package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"thesis/app/entities/login"
	"thesis/app/helpers"
	"thesis/app/services/auth"
	"thesis/config/paseto"
)

func Login(ctx *gin.Context) {
	var loginDTO *login.LoginDTO
	err := ctx.ShouldBindJSON(&loginDTO)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if !auth.Login(loginDTO.Email, loginDTO.Password) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Email or password is incorrect",
		})
		return
	}
	accessToken, _ := paseto.CreatePasetoToken(helpers.Env("SECRET_KEY_PASETO", ""), 1)
	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
	return
}
func VerifyToken(ctx *gin.Context) {
	headerToken := strings.Split(ctx.GetHeader("Authorization"), "Bearer ")
	accessToken := strings.Join(headerToken, " ")
	accessToken = strings.TrimSpace(accessToken)
	if len(accessToken) <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Token is empty",
		})
		return
	}
	key := helpers.Env("SECRET_KEY_PASETO", "")

	isVerify, errToken := paseto.VerifyToken(key, accessToken)
	if errToken != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errToken.Error(),
		})
		return
	}
	if len(isVerify) > 0 {
		ctx.JSON(http.StatusOK, gin.H{})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Token is invalid",
	})
	return
}
