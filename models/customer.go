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
	Role     string `gorm:"not null" json:"role" form:"role" valid:"in(admin|customer)~Role must be admin or customer"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) error {
	// Validate the struct using govalidator
	_, err := govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}

	hashedPass := helpers.HashPass(c.Password)
	c.Password = hashedPass

	if c.Role == "" {
		c.Role = "customer"
	} else if c.Role != "" && c.Role != "customer" {
		c.Role = "admin"
	}

	return nil
}
