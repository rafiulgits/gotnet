package gotnet

import (
	"log"
	"net/http"
)

type App struct {
	Service *Service
	Router  *Router
}

func (app *App) Run() {
	log.Printf("Server is listening %d", 5000)
	http.ListenAndServe(":5000", app.Router)
}

func (app *App) MapHandlerFunc(pattern string, handlerFn http.HandlerFunc) *App {
	app.Router.HandleFunc(pattern, handlerFn)
	return app
}

func (app *App) RegisterHandler(constructor interface{}, invoker interface{}) *App {
	app.Service.AddSingleton(constructor)
	if err := app.Service.Container().Invoke(invoker); err != nil {
		panic(err)
	}
	return app
}
