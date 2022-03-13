package gotnet

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func RequestParamInt(r *http.Request, p string) int {
	idStr := chi.URLParam(r, p)
	if idStr != "" {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return 0
		}
		return id
	}
	return 0

}

func RequestParamUInt(r *http.Request, p string) uint {
	return uint(RequestParamInt(r, p))
}

func RequestParamString(r *http.Request, p string) string {
	return chi.URLParam(r, p)
}
