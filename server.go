package main

import (
	"gitlab.com/fleet-commander/fleet-commander-backend-go/couchbase"
	"log"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/rest"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/websocket"
)

func main() {
	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Server startup ...")

	rules, err := couchbase.GetGameRules()
	if err != nil {
		log.Panicf("Can't load game rules: %+v", err)
	}
	websocket.ActiveRules = rules

	router := mux.NewRouter()
	router.HandleFunc("/websocket", websocket.ConnectionHandler)
	router.HandleFunc("/monitoring", rest.MonitoringHandler).Methods("GET")
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"x-requested-with", "authorization", "content-type", "Origin"})
	corsEnabledRouter := handlers.CORS(allowedMethods, allowedHeaders)(router)
	go websocket.KillInactiveConnections()
	log.Println(http.ListenAndServe(":8080", corsEnabledRouter))
	log.Println("Server shutdown ...")
}
