package arango

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// InsertNewPlayer inserts a new player. When either the name or the email of the
// player already exists, the functions returns a PlayerAlreadyExistsError. If the insert
// operation was successful, the given player will be extended by the ID from the
// database.
func InsertNewPlayer(player Player) error {
	passwordHash, err := player.PasswordHash()
	if err != nil {
		return errors.WithStack(err)
	}

	player.Password = passwordHash
	log.Println("insert new player:", player)

	if err := CreateDocument(&player); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// GetPlayerByEmail returns the player that matches with the given email (case insensitive)
func GetPlayerByEmail(email string) (*Player, error) {
	log.Println("get player by email:", email)

	database, err := getDatabase()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := fmt.Sprintf("FOR u IN %s FILTER LOWER(u.email) == LOWER(@email) RETURN u", playerCollectionName)
	bindings := bindingVariables{
		"email": email,
	}
	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	player := new(Player)
	_, err = cursor.ReadDocument(nil, player)
	if err != nil {
		return nil, nil
	}

	if cursor.HasMore() {
		log.Printf("ERROR: found multiple players with email %s, will use first one", email)
	}

	return player, nil
}
