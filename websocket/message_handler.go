package websocket

import (
	"encoding/json"
	"github.com/pkg/errors"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/couchbase"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/game"
	"log"
	"strings"
)

var ActiveRules game.Rules

func handleMessages(c connectedPlayer) {
	for {
		message, err := c.NextMessage()
		if err != nil {
			log.Printf("ERROR: %+v", err)
			c.SendTechnicalErrorMessage()
			c.connection.Close()
			break
		}

		switch message.Type {
		case SignIn:
			signIn(message.Payload.(*json.RawMessage), c)

		case SignUp:
			signUp(message.Payload.(*json.RawMessage), c)
		}
	}
}

func signIn(payload *json.RawMessage, c connectedPlayer) {
	login := game.Login{}
	if err := json.Unmarshal(*payload, &login); err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	playerFromDb, err := couchbase.GetPlayerByEmail(login.Email)
	if err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	hash, err := login.PasswordHash()
	if err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	if playerFromDb == nil || hash != playerFromDb.Login.Password {
		c.SendMessage(NewErrorMessage("Invalid credentials"))
		return
	}

	c.SendMessage(NewSignInMessage())
	c.SendMessage(NewGameRulesMessage())
	c.SendMessage(NewCorrectionMessage(playerFromDb.ActualResources()))
}

func signUp(payload *json.RawMessage, c connectedPlayer) {
	login := game.Login{}
	if err := json.Unmarshal(*payload, &login); err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	exists, err := playerAlreadyExists(login)
	if err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	if exists {
		c.SendMessage(NewErrorMessage("Player already exists"))
		return
	}

	err = couchbase.InsertNewPlayerWithLogin(login)
	if err != nil {
		log.Printf("ERROR: %+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	c.SendMessage(NewSignUpMessage())
}

func playerAlreadyExists(l game.Login) (bool, error) {
	playerFromDb, err := couchbase.GetPlayerByEmail(l.Email)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return playerFromDb != nil && strings.ToLower(playerFromDb.Login.Email) == strings.ToLower(l.Email) &&
		strings.ToLower(playerFromDb.Login.Name) == strings.ToLower(l.Name), nil
}
