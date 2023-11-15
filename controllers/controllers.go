package controllers

import (
	"contacts-api/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")

}

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Contacts)
}
