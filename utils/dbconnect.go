package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// GetDBconn function to return DB connection
func GetDBconn() *sql.DB {
	dbName := "AccountCrudDatabase"
	fmt.Println("conn info:", dbName)
	
	// Update the connection string for PostgreSQL
	connStr := "user=root password=root dbname=AccountCrudDatabase host=postgres port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close()

	return db
}