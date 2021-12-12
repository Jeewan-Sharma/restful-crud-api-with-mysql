package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rest_api_mysql/config"
	"rest_api_mysql/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Updateperson(w http.ResponseWriter, r *http.Request) {
	var response models.NEWResponse
	var person models.Persons

	requestBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(requestBody, &person)

	//reading dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	db := config.Connect()
	defer db.Close()

	err := db.QueryRow("UPDATE Persons set Name=? WHERE PersonID = ?", person.Name, id)
	if err != nil {
		fmt.Print("internal Server error")
	}
	response.Status = 200
	response.Message = "Updated data successfully"
	response.Data = person

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)

}
