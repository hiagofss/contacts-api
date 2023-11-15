package controllers

import (
	"contacts-api/database"
	"contacts-api/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(
		map[string]string{"status": "ok"},
	)
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

func CreateContact(w http.ResponseWriter, r *http.Request) {
	var newContact models.Contact

	json.NewDecoder(r.Body).Decode(&newContact)

	database.DB.Create(&newContact).Save(&newContact)

	json.NewEncoder(w).Encode(newContact)
}
