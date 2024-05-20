package users

import (
	"net/http"
	"todo-rest-api/internal/handlers/users"
)

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /users/{id}/", users.GetUserByID)
	router.HandleFunc("GET /users/", users.ListUsers)
}



