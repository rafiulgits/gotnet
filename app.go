package gotnet

import (
	"log"
	"net/http"
)

type App struct {
	Service *Service
	Server  *Server
}

func (app *App) Run() {
	log.Printf("Server is listening %d", 5000)
	http.ListenAndServe(":5000", app.Server.router)
}

func (app *App) MapHandler(pattern string, handlerFn http.HandlerFunc) *App {
	app.Server.router.HandleFunc(pattern, handlerFn)
	return app
}
