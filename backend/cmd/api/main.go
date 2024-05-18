package main

import (
	"todo-rest-api/config"
	"todo-rest-api/internal/application"
)



func main(){

	
	DBConfig := config.NewDBConfig("root:password@tcp(localhost:3306)/golang")
	config := config.NewConfig("8080", "development", DBConfig)
	app := application.NewApp(config, nil, nil)
	

}