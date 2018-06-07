package websocket

import (
	"encoding/json"
	"fmt"

	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/models"
)

var technicalErrorMessage = NewErrorMessage("Sorry, we have some problems with our engines, please try again later")

func handleMessages(player *connectedPlayer) {
	for {
		message, err := player.NextMessage()
		if err != nil {
			fmt.Println("ERROR: can't read message from websocket:", err)
			player.connection.Close()
			break
		}

		switch message.Type {
		case SignInType:
			signIn(message.Payload.(*json.RawMessage), player)

		case SignUpType:
			signUp(message.Payload.(*json.RawMessage), player)
		}
	}
}

func signIn(payload *json.RawMessage, player *connectedPlayer) {
	user := new(models.User)
	if err := json.Unmarshal(*payload, user); err != nil {
		fmt.Println("ERROR: can't unmarshal sign in payload:", err)
		player.SendMessage(technicalErrorMessage)
		return
	}

	userFromDb, err := arango.GetUserByEmail(user.Email)
	if err != nil {
		if err == arango.NoUserFoundError {
			player.SendMessage(NewErrorMessage("Invalid credentials"))
		} else {
			player.SendMessage(technicalErrorMessage)
		}
	} else if user.PasswordHash() != userFromDb.Password {
		player.SendMessage(NewErrorMessage("Invalid credentials"))
	} else {
		player.SendMessage(NewSignInMessage())
	}
}

func signUp(payload *json.RawMessage, player *connectedPlayer) {
	user := new(models.User)
	if err := json.Unmarshal(*payload, user); err != nil {
		fmt.Println("ERROR: can't unmarshal sign up payload:", err)
		player.SendMessage(technicalErrorMessage)
		return
	}

	if err := arango.InsertNewUser(user); err != nil {
		fmt.Println("ERROR: can't insert new user:", err)
		if err == arango.UserAlreadyExistsError {
			player.SendMessage(NewErrorMessage("User already exists"))
		} else {
			player.SendMessage(technicalErrorMessage)
		}
	} else {
		player.SendMessage(NewSignUpMessage())
	}
}
