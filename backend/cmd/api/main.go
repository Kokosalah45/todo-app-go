package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main(){

	mux := http.NewServeMux()


	mux.HandleFunc("GET /users/{id}/", func(w http.ResponseWriter, r *http.Request){
		id := r.PathValue("id")
		w.Write([]byte(id))
	})

	mux.HandleFunc("GET /users/", func(w http.ResponseWriter, r *http.Request){
		var x any

		body := r.Body

		json.NewDecoder(body).Decode(x)
		
		w.Write([]byte("LOL"))
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World"))
	})

	server := http.Server{
		Addr: ":3000",
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}