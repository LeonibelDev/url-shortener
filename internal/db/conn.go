package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func DBConnection() error {
	var err error
	Conn, err = pgxpool.New(context.Background(), os.Getenv("PSQL_CONN"))
	if err != nil {
		return fmt.Errorf("Unable to connect to database: %w", err)

	}

	fmt.Println("Connected to the database successfully")
	return nil
}

func ValidateDBExists() error {
	query := `
		CREATE TABLE IF NOT EXISTS urls (
			ID VARCHAR(10) UNIQUE PRIMARY KEY,
			link VARCHAR NOT NULL
	);`

	_, err := Conn.Exec(context.Background(), query)
	if err != nil {
		return fmt.Errorf("Unable to create table: %w", err)
	}
	return nil
}
