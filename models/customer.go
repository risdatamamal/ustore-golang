package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Customer struct {
	GormModel
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(c)
	if err != nil {
		return err
	}
	return
}
