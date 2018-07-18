package websocket

import (
	"encoding/json"
	"fmt"

	"time"

	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
)

func handleMessages(player *connectedPlayer) {
	go heardBeat(player)

	for {
		message, err := player.NextMessage()
		if err != nil {
			fmt.Printf("ERROR: can't read message from websocket \n%+v", err)
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
	user := new(arango.User)
	if err := json.Unmarshal(*payload, user); err != nil {
		fmt.Printf("ERROR: can't unmarshal sign in payload \n%+v", err)
		player.SendTechnicalErrorMessage()
		return
	}

	userFromDb, err := arango.GetUserByEmail(user.Email)
	if err != nil {
		if err == arango.NoUserFoundError {
			player.SendMessage(NewErrorMessage("Invalid credentials"))
		} else {
			player.SendTechnicalErrorMessage()
		}
	} else if user.PasswordHash() != userFromDb.Password {
		player.SendMessage(NewErrorMessage("Invalid credentials"))
	} else {
		player.SendMessage(NewSignInMessage())
	}

	if err != nil {
		fmt.Printf("%+v", err)
	}
}

func signUp(payload *json.RawMessage, player *connectedPlayer) {
	user := new(arango.User)
	if err := json.Unmarshal(*payload, user); err != nil {
		fmt.Printf("ERROR: can't unmarshal sign up payload \n%+v", err)
		player.SendTechnicalErrorMessage()
		return
	}

	user, err := arango.InsertNewUser(user)
	if err != nil {
		fmt.Printf("ERROR: can't insert new user \n%+v", err)
		if err == arango.UserAlreadyExistsError {
			player.SendMessage(NewErrorMessage("User already exists"))
		} else {
			player.SendTechnicalErrorMessage()
		}
		return
	}

	resources := new(arango.Resources)
	if err = arango.CreateDocument(resources); err != nil {
		fmt.Printf("ERROR: can't create new resources \n%+v", err)
		player.SendTechnicalErrorMessage()
		arango.RemoveDocument(user)
		return
	}

	err = arango.CreateEdge(user, resources, arango.EdgeHasResources)
	if err != nil {
		fmt.Printf("ERROR: can't create edge \n%+v", err)
		arango.RemoveDocument(user)
		arango.RemoveDocument(resources)
		player.SendTechnicalErrorMessage()
		return
	}

	player.SendMessage(NewSignUpMessage())
}

func heardBeat(player *connectedPlayer) {
	for {
		time.Sleep(5 * time.Second)
	}
}
