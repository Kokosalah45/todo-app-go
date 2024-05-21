package users

import (
	"net/http"
	"todo-rest-api/internal/handlers/users"
	"todo-rest-api/internal/middlewares"
)

func RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /users/{id}/", middlewares.WithLogging(users.GetUserByID))
	router.HandleFunc("GET /users/", middlewares.WithLogging(users.ListUsers))
}
