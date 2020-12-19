package database

import (
	"database/sql"

	_ "github.com/lib/pq" // postgres drivers
)

// ConnectWithDatabase establishes database connection
func ConnectWithDatabase() *sql.DB {
	connection := "user=postgres dbname=gocommerce password=admin host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
