package main

import (
	"time"
)

type User struct {
	id         int       `json:"id"`
	username   string    `json:"username"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}

type Expense struct {
	id         int       `json:"id"`
	user       int       `json:"user"`
	amount     float32   `json:"amount"`
	tags       []int     `json:"tags"`
	note       string    `json:"note"`
	created_at time.Time `json:"created_at"`
	updated_at time.Time `json:"updated_at"`
}

type Tag struct {
	id     int    `json:"id"`
	parent int    `json:"parent"`
	user   int    `json:"user"`
	tag    string `json:"tag"`
}

type HTTPResponse struct {
	status  int    `json:"status"`
	message string `json:"message"`
}
