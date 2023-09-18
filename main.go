package main

import (
	"fmt"
	"ustore-golang/config"
	"ustore-golang/database"
	"ustore-golang/routes"
)

func main() {
	r := routes.StartApp()
	err := database.StartDB()
	if err != nil {
		fmt.Println("Error starting database: ", err)
		return
	}
	r.Run(config.SERVER_PORT)
}
