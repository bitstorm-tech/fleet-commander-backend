package main

import (
	"fleet-commander-backend-go/resources"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Server startup ...")
	router := mux.NewRouter()
	router.HandleFunc("/users/login", corsHandler).Methods("OPTIONS")
	router.HandleFunc("/users", resources.UserCreateHandler).Methods("POST")
	router.HandleFunc("/users/login", resources.UserLoginHandler).Methods("POST")
	// loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	// corsOptions := []handlers.CORSOption{}
	// corsOptions = append(corsOptions, handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}))
	// corsOptions = append(corsOptions, handlers.AllowedOrigins([]string{"*"}))
	// corsEnabledRouter := handlers.CORS(corsOptions...)(loggedRouter)
	fmt.Println(http.ListenAndServe(":8080", router))
	fmt.Println("Server shutdown ...")
}

func corsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CORS handler!!!!")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	w.Header().Add("Access-Control-Allow-Headers", "origin, content-type, accept, authorization, x-authorization")
}
