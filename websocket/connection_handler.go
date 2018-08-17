package websocket

import (
	"fmt"
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
		fmt.Println("ERROR: can't open websocket connection:", err)
		w.WriteHeader(500)
		return
	}

	player := &connectedPlayer{
		connection: connection,
	}

	ConnectedPlayer = append(ConnectedPlayer, player)

	go handleMessages(player)
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

		fmt.Printf("Kill %v inactive connections at %v\n", killedConnectionCount, time.Now().Format(time.Stamp))

		time.Sleep(5 * time.Minute)
	}
}
