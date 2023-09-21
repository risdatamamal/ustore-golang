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
		userRouter.POST("/register", controllers.RegisterCustomer)
		userRouter.POST("/login", controllers.LoginCustomer)
		userRouter.Use(middleware.Authentication(), middleware.Authorization("userId"))
		userRouter.PUT("/:userId", controllers.UpdateCustomerData)
		userRouter.DELETE("/:userId", controllers.DeleteCustomerAccount)
	}

	return r
}
