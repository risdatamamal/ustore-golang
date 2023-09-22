package routes

import (
	"ustore-golang/controllers"
	"ustore-golang/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/api")

	{
		userRouter := v1.Group("/auth")
		{
			userRouter.POST("/register", controllers.RegisterCustomer)
			userRouter.POST("/login", controllers.LoginCustomer)
		}

		customerRouter := v1.Group("/customers")
		{
			customerRouter.Use(middleware.Authentication())
			customerRouter.GET("/", controllers.GetAllCustomers)
			customerRouter.POST("/", controllers.CreateCustomerData)
			customerRouter.Use(middleware.Authorization("customerId"))
			customerRouter.GET("/:customerId", controllers.GetDetailCustomerData)
			customerRouter.PUT("/:customerId", controllers.UpdateCustomerData)
			customerRouter.DELETE("/:customerId", controllers.DeleteCustomerData)
		}

		orderRouter := v1.Group("/orders")
		{
			orderRouter.Use(middleware.Authentication())
			orderRouter.GET("/", controllers.GetAllOrders)
			orderRouter.POST("/", controllers.CreateOrderData)
			orderRouter.Use(middleware.Authorization("orderId"))
			orderRouter.GET("/:orderId", controllers.GetDetailOrderData)
			orderRouter.PUT("/:orderId", controllers.UpdateOrderData)
			orderRouter.DELETE("/:orderId", controllers.DeleteOrderData)
		}
	}

	return r
}
