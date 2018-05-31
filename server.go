package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gitlab.com/fleet-commander/fleet-commander-backend-go/resources"
    "gitlab.com/fleet-commander/fleet-commander-backend-go/arango"
)

func main() {
	fmt.Println("Server startup ...")
	arango.Connect()
	router := mux.NewRouter()
	router.HandleFunc("/users", resources.UserCreateHandler).Methods("PUT")
	router.HandleFunc("/users/login", resources.UserLoginHandler).Methods("POST")
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"x-requested-with", "authorization", "content-type"})
	corsEnabledRouter := handlers.CORS(allowedMethods, allowedHeaders)(router)
	fmt.Println(http.ListenAndServe(":8080", corsEnabledRouter))
	fmt.Println("Server shutdown ...")
}
