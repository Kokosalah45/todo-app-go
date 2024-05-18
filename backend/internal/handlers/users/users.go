package users

import (
	"net/http"
)


func listUsers(w http.ResponseWriter , r *http.Request){}
func getUserByID(w http.ResponseWriter , r *http.Request){}

func RegisterRoutes(router *http.ServeMux){
	router.HandleFunc("GET /users/{id}/",http.HandlerFunc(getUserByID))
	router.HandleFunc("GET /users/", http.HandlerFunc(listUsers))
}

