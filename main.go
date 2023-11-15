package main

import (
	"contacts-api/database"
	"contacts-api/routes"
)

func main() {
	database.ConnectDatabase()

	routes.HandleRequest()
}
