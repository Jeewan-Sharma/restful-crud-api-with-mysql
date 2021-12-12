package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest_api_mysql/config"
	"rest_api_mysql/models"
)

func AllPersons(w http.ResponseWriter, r *http.Request) {
	var persons models.Persons
	var response models.Response
	var arrpersons []models.Persons

	db := config.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT PersonID, Name FROM persons")
	if err != nil {
		log.Print(err)
	}
	for rows.Next() {
		err = rows.Scan(&persons.Id, &persons.Name)
		if err != nil {
			log.Fatal(err.Error())
		}
		arrpersons = append(arrpersons, persons)
	}

	response.Status = 200
	response.Message = "Success"
	response.Data = arrpersons

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}
