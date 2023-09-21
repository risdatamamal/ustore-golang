package models

import (
	"ustore-golang/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Customer struct {
	GormModel
	UserName string `gorm:"not null;uniqueIndex" json:"user_name" form:"user_name" valid:"required~User name is required"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(8)" `
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email address"`
	FullName string `gorm:"null" json:"full_name" form:"full_name"`
	Role     string `gorm:"not null" json:"role" form:"role" valid:"required~Role is required"`
}

type GetAllCustomersResponse struct {
	GormModel
	UserName string `gorm:"not null;uniqueIndex" json:"user_name" form:"user_name" valid:"required~User name is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Invalid email address"`
	FullName string `gorm:"null" json:"full_name" form:"full_name"`
	Role     string `gorm:"not null" json:"role" form:"role" valid:"required~Role is required"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	hashedPass := helpers.HashPass(c.Password)

	c.Password = hashedPass
	return nil
}
