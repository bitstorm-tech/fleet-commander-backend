package main

import (
	"fleet-commander-backend-go/resources"
	"net/http"

	"fmt"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server startup ...")
	router := mux.NewRouter()
	router.HandleFunc("/users", resources.UserCreateHandler).Methods("POST")
	router.HandleFunc("/users/login", resources.UserLoginHandler).Methods("POST")
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "HEAD", "DELETE"})
	allowedHeaders := handlers.AllowedHeaders([]string{"x-requested-with", "authorization", "content-type"})
	corsEnabledRouter := handlers.CORS(allowedMethods, allowedHeaders)(router)
	fmt.Println(http.ListenAndServe(":8080", corsEnabledRouter))
	fmt.Println("Server shutdown ...")
}
