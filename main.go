package main

import (
	"log"
	"net/http"
	"rest_api_mysql/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/persons", handlers.AllPersons).Methods("GET")
	router.HandleFunc("/persons", handlers.AddPerson).Methods("POST")
	router.HandleFunc("/persons/{id}", handlers.Getperson).Methods("GET")
	router.HandleFunc("/persons/{id}", handlers.Updateperson).Methods("PUT")
	router.HandleFunc("/persons/{id}", handlers.Deleteperson).Methods("DELETE")

	log.Println("api is running")
	http.ListenAndServe(":4000", router)
}
