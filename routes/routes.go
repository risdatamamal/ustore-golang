package routes

import (
	"ustore-golang/controllers"
	"ustore-golang/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
		userRouter.Use(middleware.Authentication(), middleware.Authorization("userId"))
		userRouter.PUT("/:userId", controllers.UpdateUserData)
		userRouter.DELETE("/:userId", controllers.DeleteUserAccount)
	}

	return r
}
