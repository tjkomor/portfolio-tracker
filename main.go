package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

func main() {
	connectDB()



}

func connectDB() *sql.DB {
	setDbUser()
	connStr := fmt.Sprintf("user=%s dbname=template1 host=localhost sslmode=disable", os.Getenv("user"))
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	err = db.Ping()	
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSuccessfully connected to database!\n")

	return db
}
