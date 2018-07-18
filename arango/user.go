package arango

import (
	"crypto"
	"fmt"
)

const (
	userCollectionName = "user"
)

// User is the structure represents a user from the database
type User struct {
	Key      string `json:"_key,omitempty"`
	ID       string `json:"_id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// PasswordHash returns the user password as hex encoded SHA-512 hash string
func (u *User) PasswordHash() string {
	sha := crypto.SHA512.New()
	if _, err := sha.Write([]byte(u.Password)); err != nil {
		fmt.Println("ERROR: can't generate password hash:", err)
		return ""
	}

	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash)
}

func (User) collection() string {
	return userCollectionName
}

func (u *User) key() string {
	return u.Key
}

func (u *User) id() string {
	return u.ID
}

func (u *User) setKey(key string) {
	u.Key = key
	u.ID = u.collection() + "/" + key
}
