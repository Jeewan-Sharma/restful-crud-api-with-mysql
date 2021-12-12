package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest_api_mysql/config"
	"rest_api_mysql/models"
)

func AddPerson(w http.ResponseWriter, r *http.Request) {
	var response models.NEWResponse
	var person models.Persons

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &person)

	db := config.Connect()
	defer db.Close()
	err := db.QueryRow("INSERT INTO Persons(PersonID, Name) VALUES(?, ?)", person.Id, person.Name)
	if err != nil {

		fmt.Print("internal Server error")
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	response.Data = person
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}
