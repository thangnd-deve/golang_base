package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thesis/app/entities/users"
	"thesis/app/helpers"
	"thesis/app/services"
)

func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Implement Pleases!")
}

func CreateUser(ctx *gin.Context) {
	user := users.UserDao{}
	errorValidate := ctx.ShouldBindJSON(&user)
	if errorValidate != nil {
		errorResponse := helpers.ResponseAPIError{
			Status:  http.StatusBadRequest,
			Message: errorValidate.Error(),
			Success: false,
		}
		ctx.JSON(errorResponse.Status, errorResponse)
		return
	}
	userInfo, err := services.CreateUser(user)
	if err != nil {
		errorResponse := helpers.ResponseAPIError{
			Status:  http.StatusBadRequest,
			Message: err.Message,
			Success: false,
		}
		ctx.JSON(err.Status, errorResponse)
		return
	}
	responseSuccess := helpers.ResponseAPI{
		Status:  http.StatusCreated,
		Message: "Success",
		Success: true,
		Data:    userInfo,
	}
	ctx.JSON(responseSuccess.Status, responseSuccess)
}

func FindUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Implement Pleases!")
}

func UpdateUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Implement Pleases!")
}
func DeleteUser(ctx *gin.Context) {
	ctx.String(http.StatusNotImplemented, "Implement Pleases!")
}
