package websocket

import (
	"encoding/json"
	"fmt"

	"time"

	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
)

func handleMessages(c *playerConnection) {
	go heardBeat(c)

	for {
		message, err := c.NextMessage()
		if err != nil {
			fmt.Printf("ERROR: can't read message from websocket \n%+v", err)
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

func signIn(payload *json.RawMessage, c *playerConnection) {
	player := new(arango.Player)
	if err := json.Unmarshal(*payload, player); err != nil {
		fmt.Printf("ERROR: can't unmarshal sign in payload \n%+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	playerFromDb, err := arango.GetPlayerByEmail(player.Email)
	if err != nil {
		if err == arango.NoPlayerFoundError {
			c.SendMessage(NewErrorMessage("Invalid credentials"))
		} else {
			c.SendTechnicalErrorMessage()
		}
	} else if player.PasswordHash() != playerFromDb.Password {
		c.SendMessage(NewErrorMessage("Invalid credentials"))
	} else {
		c.SendMessage(NewSignInMessage())
	}

	if err != nil && err != arango.NoPlayerFoundError {
		fmt.Printf("%+v", err)
	}
}

func signUp(payload *json.RawMessage, c *playerConnection) {
	player := new(arango.Player)
	if err := json.Unmarshal(*payload, player); err != nil {
		fmt.Printf("ERROR: can't unmarshal sign up payload \n%+v", err)
		c.SendTechnicalErrorMessage()
		return
	}

	err := arango.InsertNewPlayer(player)
	if err != nil {
		fmt.Printf("ERROR: can't insert new player \n%+v", err)
		if err == arango.PlayerAlreadyExistsError {
			c.SendMessage(NewErrorMessage("Player already exists"))
		} else {
			c.SendTechnicalErrorMessage()
		}
		return
	}

	c.SendMessage(NewSignUpMessage())
}

func heardBeat(c *playerConnection) {
	for {
		time.Sleep(5 * time.Second)
	}
}
