package controllers

import (
	"net/http"
	"strconv"
	"ustore-golang/database"
	"ustore-golang/helpers"
	"ustore-golang/models"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(ctx *gin.Context) {
	db := database.GetDB()
	Orders := []models.Order{}

	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "Invalid page number",
		})
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": "Invalid limit number",
		})
		return
	}
	offset := (page - 1) * limit

	err = db.Offset(offset).Limit(limit).Find(&Orders).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": err.Error(),
		})
		return
	}

	var orderData []interface{}
	for _, order := range Orders {
		orderData = append(orderData, gin.H{
			"id":           order.ID,
			"customer_id":  order.CustomerID,
			"order_date":   order.OrderDate,
			"total_amount": order.TotalAmount,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":   http.StatusOK,
			"status": "success",
		},
		"data":  orderData,
		"page":  page,
		"limit": limit,
		"total": len(orderData),
	})
}

func CreateOrderData(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	// bind data from request body
	err := ctx.ShouldBind(&order)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err = db.Create(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"meta": gin.H{
			"code":    http.StatusCreated,
			"message": "success",
		},
		"data": gin.H{
			"id":           order.ID,
			"customer_id":  order.CustomerID,
			"order_date":   order.OrderDate,
			"total_amount": order.TotalAmount,
		},
	})
}

func GetDetailOrderData(ctx *gin.Context) {
	db := database.GetDB()
	order := models.Order{}

	orderId := ctx.Param("orderId")

	err := db.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": "Order not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "success",
		},
		"data": gin.H{
			"id":           order.ID,
			"customer_id":  order.CustomerID,
			"order_date":   order.OrderDate,
			"total_amount": order.TotalAmount,
		},
	})
}

func UpdateOrderData(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("orderId")
	Order := models.Order{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&Order)
	} else {
		ctx.ShouldBind(&Order)
	}

	err := db.Model(&Order).Where("id=?", id).Updates(models.Order{
		CustomerID:  Order.CustomerID,
		OrderDate:   Order.OrderDate,
		TotalAmount: Order.TotalAmount,
	}).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":    http.StatusOK,
			"message": "success",
		},
		"data": gin.H{
			"id":           id,
			"customer_id":  Order.CustomerID,
			"order_date":   Order.OrderDate,
			"total_amount": Order.TotalAmount,
		},
	})
}

func DeleteOrderData(ctx *gin.Context) {
	db := database.GetDB()
	orderID, err := strconv.Atoi(ctx.Param("orderId"))
	Order := models.Order{}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad Request",
			"message": "Invalid order ID",
		})
		return
	}

	err = db.Where("id=?", orderID).Delete(&Order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   "Not Found",
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Data has been deleted",
	})
}
