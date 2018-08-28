package rest

import (
	"net/http"

	"strconv"

	"gitlab.com/fleet-commander/fleet-commander-backend-go/websocket"
)

func MonitoringHandler(w http.ResponseWriter, r *http.Request) {
	info := "openConnections: " + strconv.Itoa(len(websocket.PlayerConnections))
	w.Write([]byte(info))
}
