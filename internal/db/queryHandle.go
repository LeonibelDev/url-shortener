package db

import (
	"context"
	"fmt"
)

func AddNewUrl(id, link string) bool {

	query := `INSERT INTO urls (ID, link) VALUES ($1, $2)`

	_, err := Conn.Exec(context.Background(), query, id, link)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("❌ Error: %s", err)
		return false
	}

	fmt.Println("✅ URL added successfully!")
	return true
}

func GetUrl(id string) string {

	var link string
	query := `SELECT link FROM urls WHERE ID = $1`
	err := Conn.QueryRow(context.Background(), query, id).Scan(&link)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return link
}
