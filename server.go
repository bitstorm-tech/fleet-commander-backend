package main

import (
	"fleet-commander-backend-go/resources"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", resources.UserPostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
