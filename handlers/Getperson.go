package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest_api_mysql/config"
	"rest_api_mysql/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Getperson(w http.ResponseWriter, r *http.Request) {
	var persons models.Persons
	var response models.NEWResponse

	//reading dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT PersonID, Name FROM persons WHERE PersonID=?", id)
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		err = rows.Scan(&persons.Id, &persons.Name)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = persons

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}
