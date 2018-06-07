package arango

import (
	"errors"
	"fmt"

	"gitlab.com/fleet-commander/fleet-commander-backend-go/models"
)

var NoUserFoundError = errors.New("no user found")
var UserAlreadyExistsError = errors.New("user already exists")

// InsertNewUser inserts a new user. When either the Username or the Email of the
// user already exists, the functions returns an error
func InsertNewUser(user *models.User) error {
	passwordHash := user.PasswordHash()
	if passwordHash == "" {
		return errors.New("can't insert user because of invalid password hash")
	}

	user.Password = passwordHash
	fmt.Println("Insert new user:", user)

	database := getDatabase()
	query := "FOR u IN users FILTER LOWER(u.email) == LOWER(@email) OR LOWER(u.username) == LOWER(@username) RETURN u"
	bindings := bindingVariables{
		"email":    user.Email,
		"username": user.Username,
	}

	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		fmt.Println("ERROR: invalid query:", err)
		return err
	}

	if cursor.HasMore() {
		fmt.Printf("WARN: user with username=%s or email=%s already exists\n", user.Username, user.Email)
		return UserAlreadyExistsError
	}

	collection, err := database.Collection(nil, "users")
	if err != nil {
		fmt.Println("ERROR: can't open collection:", err)
		return err
	}

	_, err = collection.CreateDocument(nil, user)
	if err != nil {
		fmt.Println("ERROR: can't create user:", err)
		return err
	}

	return nil
}

// GetUserByEmail returns the user that matches with the given email
func GetUserByEmail(email string) (*models.User, error) {
	fmt.Println("Get user by email:", email)

	database := getDatabase()
	query := "FOR u IN users FILTER LOWER(u.email) == LOWER(@email) RETURN u"
	bindings := bindingVariables{
		"email": email,
	}
	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		fmt.Println("ERROR: invalid query:", query)
		return nil, err
	}

	user := new(models.User)
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
