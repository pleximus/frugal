package main

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"strings"
)

var db *sql.DB

func InitDB(driver, user, password, database string) (*sql.DB, error) {
	var err error
	connectionString := fmt.Sprintf("user=%v password='%v' dbname=%v sslmode=disable", user, password, database)
	db, err = sql.Open(driver, connectionString)
	return db, err
}

// User Signup
// Returns userid on Success, 0 on Error
func AddUser(username, password string) (int, error) {
	log.Println("-- AddUser", username)

	username = strings.ToLower(username)

	// Generate Random Salt and Hash
	var randBytes = make([]byte, 16)
	rand.Read(randBytes)
	salt := hex.EncodeToString(randBytes)

	dk := pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha1.New)
	hash := hex.EncodeToString(dk)

	// Store to database
	var userid int
	err := db.QueryRow(`INSERT INTO users(username, hash, salt, created_at, updated_at) 
    VALUES($1, $2, $3, 'NOW()', 'NOW()') RETURNING id`,
		username, hash, salt).Scan(&userid)
	if err != nil {
		log.Fatal("AddUser: ", err.Error())
		return 0, err
	}
	return userid, nil
}

// User Login
// Returns userid on Success, 0 on Error
func AuthenticateUser(username, password string) (int, error) {
	log.Println("-- AuthenticateUser", username)

	username = strings.ToLower(username)

	// Store to database
	var userid int
	var hash, salt string
	err := db.QueryRow(`SELECT id, hash, salt FROM users WHERE username=$1`, username).Scan(&userid, &hash, &salt)
	if err != nil {
		log.Fatal("AuthenticateUser: ", err.Error())
		return 0, err
	}

	dk := pbkdf2.Key([]byte(password), []byte(salt), 4096, 32, sha1.New)
	generatedHash := hex.EncodeToString(dk)
	if generatedHash != hash {
		return 0, errors.New("Invalid Password")
	}
	return userid, nil
}

// Manage Tags

// Add, Update Expense
