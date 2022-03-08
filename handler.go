package gotnet

import "github.com/go-chi/chi/v5"

type IHandler interface {
	Handle(router chi.Router)
}
