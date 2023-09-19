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

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))

		//TODO: Customer
		if param == "customerId" {
			db := database.GetDB()
			customer := models.Customer{}

			err := db.Select("user_id").First(&customer, id).Error

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"code":    http.StatusNotFound,
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if userID != customer.UserID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}

			//TODO: Order
		} else if param == "orderId" {
			db := database.GetDB()
			order := models.Order{}

			err := db.Select("user_id").First(&order, id).Error

			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
					"code":    http.StatusNotFound,
					"error":   "Not Found",
					"message": "Data doesn't exist",
				})
				return
			}

			if userID != order.UserID {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"code":    http.StatusUnauthorized,
					"error":   "Unauthorized",
					"message": "You are not allowed to access this data",
				})
				return
			}
		} else {
			if userID != id {
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
