package gotnet

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	Service *Service
	Router  *chi.Mux
}

func (app *App) Run() {
	log.Printf("Server is listening %d", 5000)
	http.ListenAndServe(":5000", app.Router)
}

func (app *App) MapHandlerFunc(pattern string, handlerFn http.HandlerFunc) *App {
	app.Router.HandleFunc(pattern, handlerFn)
	return app
}

type HandlerRegistration func(router *chi.Mux, container *Container)

func (app *App) RegisterHandler(constructor interface{}, callback HandlerRegistration) *App {
	app.Service.AddSingleton(constructor)
	callback(app.Router, app.Service.Container())
	return app
}
