package websocket

import (
	"log"
	"net/http"

	"time"

	"github.com/gorilla/websocket"
)

var ConnectedPlayer = make([]*connectedPlayer, 0)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("ERROR: %+v", err)
		w.WriteHeader(500)
		return
	}

	c := &connectedPlayer{
		connection: connection,
	}

	ConnectedPlayer = append(ConnectedPlayer, c)

	go handleMessages(c)
}

func KillInactiveConnections() {
	for {
		killedConnectionCount := 0
		for _, player := range ConnectedPlayer {
			if player.lastAction.After(time.Now().Add(60 * time.Minute)) {
				player.connection.Close()
				killedConnectionCount++
			}
		}

		log.Printf("Kill %v inactive connections", killedConnectionCount)

		time.Sleep(5 * time.Minute)
	}
}
