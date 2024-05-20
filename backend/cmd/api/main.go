package main

import (
	"log"
	"net/http"
	"todo-rest-api/config"
	"todo-rest-api/internal/application"
	"todo-rest-api/internal/routes/users"
)


func main(){
	
		router := users.RegisterRoutes(http.NewServeMux())
		DBConfig := config.NewDBConfig("root:password@tcp(localhost:3306)/golang")
		config := config.NewConfig("8080", "development", DBConfig)
		app := application.NewApp(config, log.Default(), log.Default())
		app.RegisterHandler(router)
		app.Serve()

}