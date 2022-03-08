package gotnet

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	Service *Service
	Router  *Router

	beforeRun func()
}

func (app *App) Run(cfg *AppRunConfig) {
	if app.beforeRun != nil {
		app.beforeRun()
	}
	if cfg == nil {
		log.Printf("Server is listening %d", 5000)
		http.ListenAndServe(":5000", app.Router)
		return
	}
	addr := cfg.formatAddress()
	log.Printf("Server is listening %s", addr)
	http.ListenAndServe(addr, app.Router)
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

func (app *App) BeforeRun(callback func()) *App {
	app.beforeRun = callback
	return app
}

type AppRunConfig struct {
	Port int
	Host string
}

func (cfg *AppRunConfig) formatAddress() string {
	return fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
}
