package rest

import (
	"github.com/bugjoe/fleet-commander-backend/websocket"
	"net/http"
	"strconv"
)

func MonitoringHandler(w http.ResponseWriter, r *http.Request) {
	info := "openConnections: " + strconv.Itoa(len(websocket.ConnectedPlayer))
	w.Write([]byte(info))
}
