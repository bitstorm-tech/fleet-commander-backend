package websocket

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/pkg/errors"

	"time"

	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
)

func handleMessages(c *connectedPlayer) {
	go heardBeat(c)

	for {
		message, err := c.NextMessage()
		if err != nil {
			log.Printf("ERROR: %+v", err)
			c.connection.Close()
			break
		}

		switch message.Type {
		case SignInType:
			signIn(message.Payload.(*json.RawMessage), c)

		case SignUpType:
			signUp(message.Payload.(*json.RawMessage), c)
		}
	}
}

func signIn(payload *json.RawMessage, c *connectedPlayer) {
	player := new(arango.Player)
	if err := json.Unmarshal(*payload, player); err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	playerFromDb, err := arango.GetPlayerByEmail(player.Email)
	if err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	hash, err := player.PasswordHash()
	if err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	if playerFromDb == nil || hash != playerFromDb.Password {
		c.SendMessage(NewErrorMessage("Invalid credentials"))
		return
	}

	c.SendMessage(NewSignInMessage(0, 0, 0))
}

func signUp(payload *json.RawMessage, c *connectedPlayer) {
	player := arango.Player{}
	if err := json.Unmarshal(*payload, &player); err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	exists, err := playerAlreadyExists(player)
	if err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	if exists {
		c.SendMessage(NewErrorMessage("Player already exists"))
		return
	}

	err = arango.InsertNewPlayer(player)
	if err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	c.SendMessage(NewSignUpMessage())
}

func playerAlreadyExists(player arango.Player) (bool, error) {
	playerFromDb, err := arango.GetPlayerByEmail(player.Email)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return playerFromDb != nil &&
		strings.ToLower(playerFromDb.Email) == strings.ToLower(player.Email) &&
		strings.ToLower(playerFromDb.Name) == strings.ToLower(player.Name), nil
}

func heardBeat(c *connectedPlayer) {
	for {
		time.Sleep(5 * time.Second)
	}
}
