package resources

import (
	"fleet-commander-backend-go/models"
	"fleet-commander-backend-go/persistence"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	Email string
	jwt.StandardClaims
}

var signatureKey = []byte("eewOb9oQBZu8RFUmmVfkhwkdhU9l09bN")

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
		fmt.Println("ERROR: can't insert new user:", err)
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
		fmt.Println("ERROR: can't extract user from http request:", err)
		w.WriteHeader(500)
		return
	}

	userFromDb, err := persistence.GetUserByEmail(user.Email)
	if err != nil {
		fmt.Println("ERROR: no user found in database:", err)
		w.WriteHeader(401)
		return
	}

	if userFromDb.Password != user.GetPasswordHash() {
		fmt.Println("ERROR: invalid password")
		w.WriteHeader(401)
	} else {
		fmt.Println("User login is valid, create JWT")
		claims := claims{
			user.Email,
			jwt.StandardClaims{
				ExpiresAt: 15000,
			},
		}
		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
		signedJwtToken, err := jwtToken.SignedString(signatureKey)
		if err != nil {
			fmt.Println("ERROR: can't sign JWT:", err)
			w.WriteHeader(500)
		}
		fmt.Println("JWT:", signedJwtToken)
		w.Write([]byte(signedJwtToken))
	}
}
