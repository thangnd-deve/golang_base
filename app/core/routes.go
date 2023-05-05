package core

import (
	"thesis/app/controllers/auth"
	"thesis/app/controllers/users"
	"thesis/app/middleware"
)

func MapUrls() {

	loginRouter := router.Group("/login")
	{
		loginRouter.POST("", auth.Login)
	}

	router.Use(middleware.Middleware{}.Auth)
	{
		userRouter := router.Group("/users")
		{
			userRouter.GET("", users.GetUser)
			userRouter.POST("", users.CreateUser)
			userRouter.PUT("/:user_id", users.UpdateUser)
			userRouter.GET("/:user_id", users.FindUser)
			userRouter.DELETE("/:user_id", users.DeleteUser)
		}
	}
}
