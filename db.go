package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(driver, user, password, database string) (*sql.DB, error) {
	var err error
	connectionString := fmt.Sprintf("%v:%v@/%v", user, password, database)
	db, err = sql.Open(driver, connectionString)
	return db, err
}

// Signup, Login, Activate User
func AddUser() {

}

func GetUserWithTags() {

}

// Manage Tags

// Add, Update Expense
