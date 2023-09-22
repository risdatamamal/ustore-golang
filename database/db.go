package database

import (
	"fmt"
	"ustore-golang/config"
	"ustore-golang/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() error {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", config.USERNAME, config.PASSWORD, config.HOST, config.DB_NAME)
	//TODO Migration command: migrate -database "mysql://root@tcp(localhost:3306)/dbustore" -path db/migrations up
	fmt.Println(conn)
	db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})

	if err != nil {
		return err
	}

	fmt.Println("Successfully Connected to Database: ", config.DB_NAME)
	db.Debug().AutoMigrate(models.Customer{})
	return nil
}

func GetDB() *gorm.DB {
	return db
}
