package main

// User for storing user data
type User struct {
	Username string `db:"username"`
	Password string `db:"password"`
	Folder string `db:"folder"`
}