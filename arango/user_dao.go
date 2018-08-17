package arango

import (
	"fmt"

	"github.com/pkg/errors"
)

var NoUserFoundError = errors.New("no user found")
var UserAlreadyExistsError = errors.New("user already exists")

// InsertNewUser inserts a new user. When either the Username or the Email of the
// user already exists, the functions returns an UserAlreadyExistsError. If the insert
// operation was successful, the given user will be extended by the ID from the
// database.
func InsertNewUser(user *User) error {
	if user == nil {
		return errors.New("user must not be nil")
	}

	passwordHash := user.PasswordHash()
	if passwordHash == "" {
		return errors.New("can't insert user because of invalid password hash")
	}

	user.Password = passwordHash
	fmt.Println("Insert new user:", user)

	database, err := getDatabase()
	if err != nil {
		return errors.Wrapf(err, "error while inserting new user: %+v", user)
	}

	query := fmt.Sprintf("FOR u IN %s FILTER LOWER(u.email) == LOWER(@email) OR LOWER(u.username) == LOWER(@username) RETURN u", userCollectionName)
	bindings := bindingVariables{
		"email":    user.Email,
		"username": user.Username,
	}

	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return errors.Wrapf(err, "error while inserting new user: %+v", user)
	}

	if cursor.HasMore() {
		fmt.Printf("WARN: user with username=%s or email=%s already exists\n", user.Username, user.Email)
		return UserAlreadyExistsError
	}

	if err = CreateDocument(user); err != nil {
		return errors.Wrapf(err, "error while inserting new user: %+v", user)
	}

	return nil
}

// GetUserByEmail returns the user that matches with the given email
func GetUserByEmail(email string) (*User, error) {
	fmt.Println("Get user by email:", email)

	database, err := getDatabase()
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting user by email: %+v", email)
	}

	query := fmt.Sprintf("FOR u IN %s FILTER LOWER(u.email) == LOWER(@email) RETURN u", userCollectionName)
	bindings := bindingVariables{
		"email": email,
	}
	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting user by email: %+v", email)
	}

	user := new(User)
	_, err = cursor.ReadDocument(nil, user)
	if err != nil {
		fmt.Println("No user found")
		return nil, NoUserFoundError
	}

	if cursor.HasMore() {
		fmt.Printf("ERROR: found multiple users with email %s, will use first one\n", email)
	}

	return user, nil
}
