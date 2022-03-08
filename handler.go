package gotnet

import (
	"net/http"
)

type IHandler interface {
	Handle() http.Handler
}
