package routes

import (
	"ustore-golang/controllers"
	"ustore-golang/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/customers")
	{
		userRouter.POST("/register", controllers.RegisterCustomer)
		userRouter.POST("/login", controllers.LoginCustomer)
		userRouter.Use(middleware.Authentication(), middleware.Authorization("customerId"))
		userRouter.PUT("/:customerId", controllers.UpdateCustomerData)
		userRouter.DELETE("/:customerId", controllers.DeleteCustomerAccount)
	}

	return r
}
