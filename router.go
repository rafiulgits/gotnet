package gotnet

import (
	"github.com/go-chi/chi/v5"
)

type Router struct {
	*chi.Mux
}

func NewRouter() *Router {
	return &Router{
		chi.NewRouter(),
	}
}
