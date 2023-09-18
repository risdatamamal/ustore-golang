package main

import (
	"fmt"

	"github.com/risdatamamal/ustore-golang/config"
	"github.com/risdatamamal/ustore-golang/database"
	"github.com/risdatamamal/ustore-golang/router"
)

func main() {
	r := router.StartApp()
	err := database.StartDB()
	if err != nil {
		fmt.Println("Error starting database: ", err)
		return
	}
	r.Run(config.SERVER_PORT)
}
