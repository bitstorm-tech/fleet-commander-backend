package arango

import (
	"fmt"

	"github.com/pkg/errors"
)

var NoUserFoundError = errors.New("no user found")
var UserAlreadyExistsError = errors.New("user already exists")

// InsertNewUser inserts a new user. When either the Username or the Email of the
// user already exists, the functions returns an error
func InsertNewUser(user *User) (*User, error) {
	passwordHash := user.PasswordHash()
	if passwordHash == "" {
		return nil, errors.New("can't insert user because of invalid password hash")
	}

	user.Password = passwordHash
	fmt.Println("Insert new user:", user)

	database, err := GetDatabase()
	if err != nil {
		return nil, errors.Wrapf(err, "error while inserting new user: %v\n", user)
	}

	query := fmt.Sprintf("FOR u IN %s FILTER LOWER(u.email) == LOWER(@email) OR LOWER(u.username) == LOWER(@username) RETURN u", CollectionUser)
	bindings := bindingVariables{
		"email":    user.Email,
		"username": user.Username,
	}

	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return nil, errors.Wrapf(err, "error while inserting new user: %v\n", user)
	}

	if cursor.HasMore() {
		fmt.Printf("WARN: user with username=%s or email=%s already exists\n", user.Username, user.Email)
		return nil, UserAlreadyExistsError
	}

	if err = CreateDocument(user); err != nil {
		return nil, errors.Wrapf(err, "error while inserting new user: %v\n", user)
	}

	return user, nil
}

// GetUserByEmail returns the user that matches with the given email
func GetUserByEmail(email string) (*User, error) {
	fmt.Println("Get user by email:", email)

	database, err := GetDatabase()
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting user by email: %v\n", email)
	}

	query := fmt.Sprintf("FOR u IN %s FILTER LOWER(u.email) == LOWER(@email) RETURN u", CollectionUser)
	bindings := bindingVariables{
		"email": email,
	}
	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting user by email: %v\n", email)
	}

	user := new(User)
	_, err = cursor.ReadDocument(nil, user)
	if err != nil {
		fmt.Println("ERROR: no user found with email:", email)
		return nil, NoUserFoundError
	}

	if cursor.HasMore() {
		fmt.Printf("ERROR: found multiple users with email %s, will use first one\n", email)
	}

	return user, nil
}
