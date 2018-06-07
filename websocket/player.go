package websocket

import (
	"time"

	"encoding/json"

	"fmt"

	"github.com/gorilla/websocket"
)

type connectedPlayer struct {
	id         string
	connection *websocket.Conn
	lastAction time.Time
	created    time.Time
}

func (player *connectedPlayer) NextMessage() (*Message, error) {
	message := &Message{
		Payload: &json.RawMessage{},
	}

	if err := player.connection.ReadJSON(message); err != nil {
		return nil, err
	}

	return message, nil
}

func (player *connectedPlayer) SendMessage(message *Message) error {
	if err := player.connection.WriteJSON(message); err != nil {
		fmt.Println("ERROR: can't write sign in answer:", err)
		return err
	}

	return nil
}
