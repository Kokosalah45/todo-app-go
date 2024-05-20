package main

import (
	"log"
	"net/http"
	"todo-rest-api/config"
	"todo-rest-api/internal/application"
	"todo-rest-api/internal/routes/users"
)


func main(){
		DBConfig := config.NewDBConfig("root:password@tcp(localhost:3306)/golang")
		config := config.NewConfig("8080", "development", DBConfig)

		router := http.NewServeMux()
		users.RegisterRoutes(router)
		

		app := application.NewApp(config, log.Default(), log.Default())
		app.RegisterRouter(router)
		app.Serve()


}