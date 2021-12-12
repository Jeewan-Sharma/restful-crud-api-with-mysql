package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:Sharma123@tcp(localhost:3306)/crud_api")
	if err != nil {
		panic(err.Error())
	}
	return db
}
