package main

import (
	"database/sql"
	"example.com/bookstore/dataobjects"
	"fmt"
	"log"
)

func main() {

	var db *sql.DB

	db, err := dataobjects.GetConnection(db)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected!")

}
