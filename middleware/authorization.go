package middleware

import (
	"net/http"
	"strconv"
	"ustore-golang/database"
	"ustore-golang/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authorization(param string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param(param))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"error":   "Bad request",
				"message": "Invalid ID",
			})
			return
		}

		customerData := ctx.MustGet("customerData").(jwt.MapClaims)
		customerID := int(customerData["id"].(float64))

		//TODO: Order
		if param == "orderId" {
			db := database.GetDB()
			order := models.Order{}

			err := db.Select("customer_id").First(&order, id).Error

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"code":    http.StatusNotFound,
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if customerID != order.CustomerID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}
		} else {
			if customerID != id {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}
		}

		ctx.Next()
	}
}
