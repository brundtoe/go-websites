package main

import (
	"brundtoe/bookstore/dataobjects"
	"database/sql"
	"fmt"
	"log"
)

func main() {

	var db *sql.DB

	// Fungerer kun når MariaBD er installeret på hosten hvor programmet udføres

	db, err := dataobjects.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected!")

}
