package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thesis/app/entities/users"
	"thesis/app/helpers"
	"thesis/app/services"
)

func GetUser(ctx *gin.Context) {
	var userDao []users.UserDTO
	userEntity := services.UserService{}.GetAllUser()
	for i := 0; i < len(userEntity); i++ {
		userDaoLocal := users.UserDTO{
			Name:  userEntity[i].Name,
			Email: userEntity[i].Email,
		}
		userDao = append(userDao, userDaoLocal)
	}

	ctx.JSON(http.StatusOK, helpers.ResponseAPI{
		Success: true,
		Status:  http.StatusOK,
		Message: "Success",
		Data:    userDao,
	})
	return
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
	userInfo, err := services.UserService{}.CreateUser(user)
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
