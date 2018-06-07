package models

import (
	"crypto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// User is the structure represents a user from the database
type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// GetPasswordHash returns the user password as hex encoded SHA-512 hash string
func (user *User) PasswordHash() string {
	sha := crypto.SHA512.New()
	if _, err := sha.Write([]byte(user.Password)); err != nil {
		fmt.Println("ERROR: can't generate password hash:", err)
		return ""
	}

	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

// UserFromRequest extracts the user from a request
func UserFromRequest(r *http.Request) (*User, error) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ERROR: can't read request body", err)
		return nil, err
	}

	user := new(User)
	err = json.Unmarshal(b, user)
	if err != nil {
		fmt.Println("ERROR: can't unmarshal request body:", string(b), err)
		return nil, err
	}

	return user, nil
}
