package resources

import (
	"encoding/json"
	"fleet-commander-backend-go/models"
	"fmt"
	"io/ioutil"
	"net/http"
    "fleet-commander-backend-go/persistence"
)

// UserPostHandler handles the POST requests
func UserPostHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("ERROR: Can't read request body")
		w.WriteHeader(500)
	}

	user := new(models.User)
	if json.Unmarshal(b, user) != nil {
		fmt.Println("ERROR: Can't unmarshal request body:", string(b))
	}

	persistence.UpsertUser(user)
}
