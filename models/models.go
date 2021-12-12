package models

type Persons struct {
	Id   int32  `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}

//for showing response
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Persons
}

type NEWResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Persons
}
