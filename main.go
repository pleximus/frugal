package main

import (
	"log"
)

func main() {

	// Init DB
	db, err := InitDB("postgres", "postgres", "root", "frugal")
	if err != nil {
		log.Fatal("InitDB ", err.Error())
		return
	}
	defer db.Close()

	// Start HTTP Server
	err = StartServer("0.0.0.0", "8001")
	if err != nil {
		log.Fatal("StartServer ", err.Error())
		return
	}

}
