package application

import (
	"log"
	"net/http"
	"todo-rest-api/config"
)

const version = "1.0.0"


type application struct {
	config *config.Config
	infoLog *log.Logger
	errorLog *log.Logger
	version string
}

func NewApp(config *config.Config, infoLog *log.Logger, errorLog *log.Logger) *application {
	return &application{
		config: config,
		infoLog: infoLog,
		errorLog: errorLog,
		version: version,
	}
}

func (app *application) serve() {
	srv := &http.Server{
		Addr: ":" + app.config.Port,
		Handler: app.routes(),
		ErrorLog: app.errorLog,
	}

	app.infoLog.Printf("Starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)

}