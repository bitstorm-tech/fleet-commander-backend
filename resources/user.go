package resources

import (
	"fleet-commander-backend-go/models"
	"fleet-commander-backend-go/persistence"
	"fmt"
	"net/http"
	"strings"
)

// UserCreateHandler creates a new user in the database
func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.UserFromRequest(r)
	if err != nil {
		fmt.Println("ERROR: Can't extract user from http request:", err)
		w.WriteHeader(500)
		return
	}

	err = persistence.InsertNewUser(user)
	if err != nil {
		fmt.Println("ERROR: Can't insert new user:", err)
		if strings.Contains(err.Error(), "already exists") {
			w.WriteHeader(403)
		} else {
			w.WriteHeader(500)
		}
	}
}

// UserLoginHandler checks if the submitted user is valid and returns
// a JWT token if that is the case
func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.UserFromRequest(r)
	if err != nil {
		fmt.Println("ERROR: Can't extract user from http request:", err)
		w.WriteHeader(500)
		return
	}

	userFromDb, err := persistence.GetUserByEmail(user.Email)
	if err != nil {
		fmt.Println("ERROR: No user found in database:", err)
		w.WriteHeader(401)
		return
	}

	if userFromDb.Password != user.GetPasswordHash() {
		fmt.Println("ERROR: Invalid password")
		w.WriteHeader(401)
	} else {
		fmt.Println("User login is valid, create JWT")
		w.Write([]byte("ThisIsMySuperDuperCoolJWT!!!"))
	}
}
