package gotnet

import "github.com/go-chi/chi/v5"

type AppBuilder struct {
	Service *Service
}

func NewAppBuilder() *AppBuilder {
	return &AppBuilder{
		Service: newService(),
	}
}

func (builder *AppBuilder) Build() *App {
	app := &App{
		Service: builder.Service,
		Router:  chi.NewRouter(),
	}
	return app
}
