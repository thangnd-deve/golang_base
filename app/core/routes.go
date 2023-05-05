package core

import (
	"thesis/app/controllers/auth"
	"thesis/app/controllers/users"
)

func MapUrls() {
	userRouter := router.Group("/users")
	{
		userRouter.GET("", users.GetUser)
		userRouter.POST("", users.CreateUser)
		userRouter.PUT("/:user_id", users.UpdateUser)
		userRouter.GET("/:user_id", users.FindUser)
		userRouter.DELETE("/:user_id", users.DeleteUser)
	}

	loginRouter := router.Group("/login")
	{
		loginRouter.POST("", auth.Login)
	}

	checkRouter := router.Group("verify-token")
	{
		checkRouter.POST("", auth.VerifyToken)
	}
}
