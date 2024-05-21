package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-rest-api/config"
	"todo-rest-api/internal/application"
	"todo-rest-api/internal/routes/users"
	"todo-rest-api/internal/storage"
)


func main(){

		port := os.Getenv("API_PORT")
		environment := os.Getenv("ENV")

		if port == "" {
			port = "8080"
		}

		if environment == "" {
			environment = "development"
		}

		dbName := os.Getenv("POSTGRES_DB")
		dbUser := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		dbHost := os.Getenv("POSTGRES_HOST")
		
		
		store , err := storage.NewPostgresStore(dbUser, dbName, dbPassword , dbHost)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(store)
		
		config := config.NewConfig(port, environment)

		router := http.NewServeMux()

		users.RegisterRoutes(router)
	
		app := application.NewApp(config, log.Default(), log.Default())
		app.RegisterRouter(router)
		app.Serve()


}