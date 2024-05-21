package main

import (
	"log"
	"net/http"
	"os"
	"todo-rest-api/config"
	"todo-rest-api/internal/application"
	"todo-rest-api/internal/routes/users"
)


func main(){
		port := os.Getenv("API_PORT")

		if port == "" {
			port = "8080"
		}
		
		DBConfig := config.NewDBConfig("root:password@tcp(localhost:3306)/golang")
		config := config.NewConfig(port, "development", DBConfig)

		router := http.NewServeMux()

		users.RegisterRoutes(router)
	
		app := application.NewApp(config, log.Default(), log.Default())
		app.RegisterRouter(router)
		app.Serve()


}