package users

import (
	"net/http"
	"todo-rest-api/internal/handlers/users"
)

func RegisterRoutes(router *http.ServeMux) *http.ServeMux{
	router.HandleFunc("GET /users/{id}/",http.HandlerFunc(users.GetUserByID))
	router.HandleFunc("GET /users/", http.HandlerFunc(users.ListUsers))
	return router
}

