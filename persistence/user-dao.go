package persistence

import (
	"fleet-commander-backend-go/models"
	"fmt"
)

// UpsertUser upserts the given user into the database
func UpsertUser(user *models.User) {
	fmt.Printf("Upsert user: %v\n", user)

	c := newArangoClient()
	d, err := c.Database(nil, "fleet-commander")
	if err != nil {
		fmt.Println("ERROR: Can't open database", err)
		return
	}

	col, err := d.Collection(nil, "users")
	if err != nil {
		fmt.Println("ERROR: Can't open collection", err)
		return
	}

	col.CreateDocument(nil, user)
}
