package database

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func SetDB(database *sql.DB) {
	Db = database
	fmt.Println("Successfully set connection to database")
}
