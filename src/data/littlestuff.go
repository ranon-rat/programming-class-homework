package data

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func getConnection() *sql.DB {

	database, _ := sql.Open("sqlite3", "./sql/database.db")
	return database
}
