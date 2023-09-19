package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Order struct {
	GormModel
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	_, err = govalidator.ValidateStruct(o)
	if err != nil {
		return err
	}
	return
}
