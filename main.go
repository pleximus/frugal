package main

import (
	"log"
)

func main() {
	// Init DB
	db, err := InitDB("postgres", "postgres", "", "frugal")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	// Start HTTP Server
	err = StartServer("0.0.0.0", "8001")
	if err != nil {
		log.Fatal(err)
		return
	}
}
