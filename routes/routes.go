package routes

import (
	"contacts-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/contacts", controllers.GetContacts).Methods("GET")
	r.HandleFunc("/contacts/{id}", controllers.GetContactById).Methods("GET")
	r.HandleFunc("/contacts", controllers.CreateContact).Methods("POST")
	r.HandleFunc("/contacts/{id}", controllers.DeleteContact).Methods("DELETE")
	r.HandleFunc("/contacts/{id}", controllers.EditContact).Methods("PUT")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
