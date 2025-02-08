package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func ValidateDBExists() bool {
	db := DBConnection()
	defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS Urls (
		ID TEXT PRIMARY KEY UNIQUE,
		Link TEXT
	);
	`
	_, err := db.Exec(createTable)
	if err != nil {
		fmt.Println("Error: ", err)
		return false
	}

	return true
}
