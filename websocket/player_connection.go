package websocket

import (
	"time"

	"encoding/json"

	"fmt"

	"github.com/gorilla/websocket"
)

type playerConnection struct {
	id         string
	connection *websocket.Conn
	lastAction time.Time
	created    time.Time
}

func (c *playerConnection) NextMessage() (*Message, error) {
	message := &Message{
		Payload: &json.RawMessage{},
	}

	if err := c.connection.ReadJSON(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (c *playerConnection) SendMessage(message *Message) error {
	if err := c.connection.WriteJSON(message); err != nil {
		fmt.Println("ERROR: can't write sign in answer:", err)
		return err
	}

	return nil
}

func (c *playerConnection) SendTechnicalErrorMessage() {
	c.SendMessage(NewErrorMessage("Sorry, we have some problems with our engines, please try again later"))
}
