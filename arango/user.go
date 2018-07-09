package arango

import (
	"crypto"
	"fmt"
)

// User is the structure represents a user from the database
type User struct {
	Key_     string `json:"_key"`
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

func (User) Collection() string {
	return CollectionUser
}

func (user *User) ID() string {
	if len(user.Key()) == 0 {
		return ""
	}

	return user.Collection() + "/" + user.Key()
}

func (user *User) Key() string {
	return user.Key_
}
