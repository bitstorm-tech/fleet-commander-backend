package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/websocket"
)

func main() {
	fmt.Println("Server startup ...")
	arango.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/websocket", websocket.ConnectionHandler)
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"x-requested-with", "authorization", "content-type", "Origin"})
	corsEnabledRouter := handlers.CORS(allowedMethods, allowedHeaders)(router)
	fmt.Println(http.ListenAndServe(":8080", corsEnabledRouter))
	fmt.Println("Server shutdown ...")
}
