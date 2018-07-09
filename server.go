package main

import (
	"fmt"

	"net/http"

	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/rest"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/websocket"
)

func main() {
	fmt.Println("Server startup ...")
	err := arango.Setup()
	if err != nil {
		fmt.Printf("error while starting server %+v\n", err)
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/websocket", websocket.ConnectionHandler)
	router.HandleFunc("/monitoring", rest.MonitoringHandler).Methods("GET")
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"x-requested-with", "authorization", "content-type", "Origin"})
	corsEnabledRouter := handlers.CORS(allowedMethods, allowedHeaders)(router)
	go websocket.KillInactiveConnections()
	fmt.Println(http.ListenAndServe(":8080", corsEnabledRouter))
	fmt.Println("Server shutdown ...")
}
