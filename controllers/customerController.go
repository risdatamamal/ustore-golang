package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"ustore-golang/database"
	"ustore-golang/helpers"
	"ustore-golang/models"

	"github.com/gin-gonic/gin"
)

func RegisterCustomer(ctx *gin.Context) {
	db := database.GetDB()
	Customer := models.Customer{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&Customer)
	} else {
		ctx.ShouldBind(&Customer)
	}

	err := db.Debug().Create(&Customer).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"meta": gin.H{
			"code":   http.StatusCreated,
			"status": "success",
		},
		"data": gin.H{
			"id":         Customer.ID,
			"user_name":  Customer.UserName,
			"email":      Customer.Email,
			"role":       Customer.Role,
			"created_at": Customer.CreatedAt,
		},
	})
}

func LoginCustomer(ctx *gin.Context) {
	db := database.GetDB()
	Customer := models.Customer{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&Customer)
	} else {
		ctx.ShouldBind(&Customer)
	}

	password := Customer.Password
	err := db.Debug().Where("email = ? ", Customer.Email).Take(&Customer).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": fmt.Sprintf("Email Not registered :%s", err.Error()),
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(Customer.Password), []byte(password))
	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"error":   "Unauthorized",
			"message": "Wrong password",
		})
		return
	}

	token, err := helpers.GenerateToken(Customer.ID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": fmt.Sprintf("Error generating token :%s", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":   http.StatusOK,
			"status": "success",
		},
		"data": gin.H{
			"token":     token,
			"id":        Customer.ID,
			"user_name": Customer.UserName,
			"email":     Customer.Email,
		},
	})
}

func GetAllCustomers(ctx *gin.Context) {
	db := database.GetDB()
	Customers := []models.Customer{}

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

	err = db.Offset(offset).Limit(limit).Find(&Customers).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"error":   "Internal server error",
			"message": err.Error(),
		})
		return
	}

	var customerData []interface{}
	for _, customer := range Customers {
		customerData = append(customerData, gin.H{
			"id":        customer.ID,
			"user_name": customer.UserName,
			"email":     customer.Email,
			"role":      customer.Role,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":   http.StatusOK,
			"status": "success",
		},
		"data":  customerData,
		"page":  page,
		"limit": limit,
		"total": len(customerData),
	})
}

func CreateCustomerData(ctx *gin.Context) {
	db := database.GetDB()
	Customer := models.Customer{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&Customer)
	} else {
		ctx.ShouldBind(&Customer)
	}

	err := db.Debug().Create(&Customer).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"meta": gin.H{
			"code":   http.StatusCreated,
			"status": "success",
		},
		"data": gin.H{
			"id":         Customer.ID,
			"user_name":  Customer.UserName,
			"email":      Customer.Email,
			"created_at": Customer.CreatedAt,
		},
	})
}

func GetDetailCustomerData(ctx *gin.Context) {
	db := database.GetDB()
	customerID, err := strconv.Atoi(ctx.Param("customerId"))
	Customer := models.Customer{}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad Request",
			"message": "Invalid customer ID",
		})
		return
	}

	err = db.Where("id=?", customerID).First(&Customer).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"error":   "Not Found",
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"meta": gin.H{
			"code":   http.StatusOK,
			"status": "success",
		},
		"data": gin.H{
			"id":        Customer.ID,
			"user_name": Customer.UserName,
			"email":     Customer.Email,
			"full_name": Customer.FullName,
			"role":      Customer.Role,
		},
	})
}

func UpdateCustomerData(ctx *gin.Context) {
	db := database.GetDB()
	id := ctx.Param("customerId")
	Customer := models.Customer{}

	reqHeaders := helpers.GetRequestHeaders(ctx)
	if reqHeaders.ContentType == "application/json" {
		ctx.ShouldBindJSON(&Customer)
	} else {
		ctx.ShouldBind(&Customer)
	}

	err := db.Model(&Customer).Where("id=?", id).Updates(models.Customer{
		Email:    Customer.Email,
		UserName: Customer.UserName,
		FullName: Customer.FullName,
		Role:     Customer.Role,
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
			"code":   http.StatusOK,
			"status": "success",
		},
		"data": gin.H{
			"id":         id,
			"email":      Customer.Email,
			"user_name":  Customer.UserName,
			"updated_at": Customer.UpdatedAt,
		},
	})
}

func DeleteCustomerData(ctx *gin.Context) {
	db := database.GetDB()
	customerID, err := strconv.Atoi(ctx.Param("customerId"))
	Customer := models.Customer{}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"error":   "Bad Request",
			"message": "Invalid customer ID",
		})
		return
	}

	err = db.Where("id=?", customerID).Delete(&Customer).Error
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
