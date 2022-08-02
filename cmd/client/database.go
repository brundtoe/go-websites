package main

import (
	"database/sql"
	"fmt"
	"github.com/brundtoe/bookstore/dataobjects"
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
