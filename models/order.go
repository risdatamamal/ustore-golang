package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	GormModel
	CustomerID  int `json:"customer_id"`
	Customer    *Customer
	OrderDate   string `gorm:"not null" json:"order_date" form:"order_date" valid:"required~Order date is required"`
	TotalAmount int    `gorm:"not null" json:"total_amount" form:"total_amount" valid:"required~Total amount is required"`
}

type GetAllOrdersResponse struct {
	GormModel
	CustomerID int `json:"customer_id"`
	Customer   struct {
		UserName string `json:"user_name"`
		Email    string `json:"email"`
	}
	OrderDate   string `gorm:"not null" json:"order_date" form:"order_date" valid:"required~Order date is required"`
	TotalAmount int    `gorm:"not null" json:"total_amount" form:"total_amount" valid:"required~Total amount is required"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(o)
	if err != nil {
		return err
	}
	return nil
}
