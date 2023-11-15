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
	r.HandleFunc("/contacts", controllers.GetAllContacts)
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
