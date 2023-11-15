package main

import (
	"contacts-api/models"
	"contacts-api/routes"
)

func main() {
	models.Contacts = []models.Contact{
		{Id: 1, Name: "John Doe", Phone: "1234567890", Email: "john.doe@test.com"},
		{Id: 2, Name: "Jane Doe", Phone: "1234567890", Email: "jane.doe@test.com"},
	}

	routes.HandleRequest()
}
