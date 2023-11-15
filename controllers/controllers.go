package controllers

import (
	"contacts-api/database"
	"contacts-api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")

}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	var c []models.Contact
	database.DB.Find(&c)

	json.NewEncoder(w).Encode(c)
}

func GetContactById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	c := models.Contact{}

	database.DB.First(&c, "id = ?", id)

	json.NewEncoder(w).Encode(c)
}
