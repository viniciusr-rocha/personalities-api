package routes

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"personalities-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
