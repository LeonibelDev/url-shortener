package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func DBConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "internal/db/urls.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return db
}
