package middleware

import (
	"net/http"
	"strconv"

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

		ctx.Next()
	}
}
