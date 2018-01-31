package models

// User is the structure represents a user from the database
type User struct {
	Username string
	Password string
	Email    string
	JWT      string
}
