package gotnet

import "github.com/go-chi/chi"

type Router struct {
	*chi.Mux
}

func newRouter() *Router {
	return &Router{chi.NewRouter()}
}
