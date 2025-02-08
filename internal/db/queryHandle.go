package db

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddNewUrl(id, link string) bool {
	db := DBConnection()
	defer db.Close()

	query := `INSERT INTO Urls (ID, Link) VALUES (?, ?)`
	_, err := db.Exec(query, id, link)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("❌ Error: %s", err)
		return false
	}

	fmt.Println("✅ URL added successfully!")
	return true
}

func GetUrl(id string) string {
	db := DBConnection()
	defer db.Close()

	var link string
	query := `SELECT Link FROM Urls WHERE ID = ?`
	err := db.QueryRow(query, id).Scan(&link)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return link
}
