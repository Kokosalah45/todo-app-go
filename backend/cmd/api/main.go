package main

import (
	"log"
	"net/http"
	"todo-rest-api/internal/handlers/users"
)

func main(){

	mux := http.NewServeMux()

	users.RegisterUsersRoutes(mux)

	server := http.Server{
		Addr: ":3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}