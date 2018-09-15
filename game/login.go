package game

import (
	"crypto"
	"fmt"
	"github.com/pkg/errors"
)

type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// PasswordHash returns the players password as hex encoded SHA-512 hash string
func (l Login) PasswordHash() (string, error) {
	sha := crypto.SHA512.New()
	if _, err := sha.Write([]byte(l.Password)); err != nil {
		return "", errors.WithStack(err)
	}

	hash := sha.Sum(nil)

	return fmt.Sprintf("%x", hash), nil
}
