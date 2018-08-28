package arango

import (
	"fmt"

	"github.com/pkg/errors"
)

var NoPlayerFoundError = errors.New("no player found")
var PlayerAlreadyExistsError = errors.New("player already exists")

// InsertNewPlayer inserts a new player. When either the name or the email of the
// player already exists, the functions returns a PlayerAlreadyExistsError. If the insert
// operation was successful, the given player will be extended by the ID from the
// database.
func InsertNewPlayer(player *Player) error {
	if player == nil {
		return errors.New("player must not be nil")
	}

	passwordHash := player.PasswordHash()
	if passwordHash == "" {
		return errors.New("can't insert player because of invalid password hash")
	}

	player.Password = passwordHash
	fmt.Println("Insert new player:", player)

	database, err := getDatabase()
	if err != nil {
		return errors.Wrapf(err, "error while inserting new player: %+v", player)
	}

	query := fmt.Sprintf("FOR u IN %s FILTER LOWER(u.email) == LOWER(@email) OR LOWER(u.name) == LOWER(@name) RETURN u", playerCollectionName)
	bindings := bindingVariables{
		"email": player.Email,
		"name":  player.Name,
	}

	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return errors.Wrapf(err, "error while inserting new player: %+v", player)
	}

	if cursor.HasMore() {
		fmt.Printf("WARN: player with name=%s or email=%s already exists\n", player.Name, player.Email)
		return PlayerAlreadyExistsError
	}

	if err = CreateDocument(player); err != nil {
		return errors.Wrapf(err, "error while inserting new player: %+v", player)
	}

	return nil
}

// GetPlayerByEmail returns the player that matches with the given email
func GetPlayerByEmail(email string) (*Player, error) {
	fmt.Println("Get player by email:", email)

	database, err := getDatabase()
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting player by email: %+v", email)
	}

	query := fmt.Sprintf("FOR u IN %s FILTER LOWER(u.email) == LOWER(@email) RETURN u", playerCollectionName)
	bindings := bindingVariables{
		"email": email,
	}
	cursor, err := database.Query(nil, query, bindings)
	if err != nil {
		return nil, errors.Wrapf(err, "error while getting player by email: %+v", email)
	}

	player := new(Player)
	_, err = cursor.ReadDocument(nil, player)
	if err != nil {
		fmt.Println("No player found")
		return nil, NoPlayerFoundError
	}

	if cursor.HasMore() {
		fmt.Printf("ERROR: found multiple players with email %s, will use first one\n", email)
	}

	return player, nil
}
