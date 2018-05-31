package resources

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/models"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
)

type claims struct {
	Email string
	jwt.StandardClaims
}

var expireTime = int64(1000 * 60 * 60 * 24)
var signatureKey = []byte("eewOb9oQBZu8RFUmmVfkhwkdhU9l09bN")

// UserCreateHandler creates a new user in the database
func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.UserFromRequest(r)
	if err != nil {
		fmt.Println("ERROR: Can't extract user from http request:", err)
		w.WriteHeader(500)
		return
	}

	err = arango.InsertNewUser(user)
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

	userFromDb, err := arango.GetUserByEmail(user.Email)
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
				ExpiresAt: time.Now().Unix()*1000 + expireTime,
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
