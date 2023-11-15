package controllers

import (
	"contacts-api/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")

}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Contacts)
}

func GetContactById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	idConverted, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, "Invalid Id")
	}

	for _, contact := range models.Contacts {
		if contact.Id == idConverted {
			json.NewEncoder(w).Encode(contact)
		}
	}
}
