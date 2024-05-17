package users

import (
	"net/http"
	"todo-rest-api/internal/middlewares"
)




func listUsers(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("List of users"))
}
func getUserByID(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("User by ID"))
}

 
func RegisterUsersRoutes(mux *http.ServeMux){
	mux.HandleFunc("GET /users/{id}/", middlewares.WithLogging(getUserByID))
	mux.HandleFunc("GET /users/", middlewares.WithLogging(listUsers))
}