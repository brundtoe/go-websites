package main

import (
	"brundtoe/bookstore/dataobjects"
	"database/sql"
	"fmt"
	"log"
)

func main() {

	var db *sql.DB

	db, err := dataobjects.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected!")

}
