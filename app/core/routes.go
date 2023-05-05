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
			userRouter.PUT("/:userId", users.UpdateUser)
			userRouter.DELETE("/:userId", users.DeleteUser)
		}
	}
}
