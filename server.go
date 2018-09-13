package main

import (
	"github.com/bugjoe/fleet-commander-backend/couchbase"
	"github.com/bugjoe/fleet-commander-backend/rest"
	"github.com/bugjoe/fleet-commander-backend/websocket"
	"log"

	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
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
