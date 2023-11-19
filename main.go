package main

import (
	"fmt"
	"github.com/fayca121/stock-pg/config"
	"github.com/fayca121/stock-pg/routes"
	"log"
)

func main() {

	db := config.ConnectDB()
	defer config.DisconnectDB(db)

	route := routes.Routes(db)

	if err := route.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server running on 8080")
}
