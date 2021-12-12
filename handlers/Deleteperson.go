package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest_api_mysql/config"
	"rest_api_mysql/models"
	"strconv"

	"github.com/gorilla/mux"
)

func Deleteperson(w http.ResponseWriter, r *http.Request) {
	var response models.NEWResponse
	var person models.Persons

	//reading dynamic id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	db := config.Connect()
	defer db.Close()

	err := db.QueryRow("DELETE FROM Persons WHERE PersonID = ?", id)
	if err != nil {
		fmt.Print("internal Server error")
	}
	response.Status = 200
	response.Message = "Deleted data successfully"
	response.Data = person

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
