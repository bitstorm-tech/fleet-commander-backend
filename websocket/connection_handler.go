package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("ERROR: can't open websocket connection:", err)
		w.WriteHeader(500)
		return
	}

	player := &connectedPlayer{
		connection: connection,
	}

	go handleMessages(player)
}
