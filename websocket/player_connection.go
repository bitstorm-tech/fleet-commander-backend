package websocket

import (
	"github.com/pkg/errors"
	"log"
	"time"

	"encoding/json"

	"github.com/gorilla/websocket"
)

type connectedPlayer struct {
	id         string
	connection *websocket.Conn
	lastAction time.Time
	created    time.Time
}

func (c connectedPlayer) NextMessage() (Message, error) {
	message := Message{
		Payload: &json.RawMessage{},
	}

	if err := c.connection.ReadJSON(&message); err != nil {
		return Message{}, errors.WithStack(err)
	}

	return message, nil
}

func (c connectedPlayer) SendMessage(message Message) error {
	if err := c.connection.WriteJSON(message); err != nil {
		log.Printf("ERROR: %+v", err)
		return err
	}

	return nil
}

func (c connectedPlayer) SendTechnicalErrorMessage() {
	c.SendMessage(NewErrorMessage("Sorry, we have some problems with our engines, please try again later"))
}
