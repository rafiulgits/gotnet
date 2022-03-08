package gotnet

import (
	"encoding/json"
	"net/http"
)

//Ok : HTTP 200 response
func Ok(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(d)
}

//Created : HTTP 201 response
func Created(w http.ResponseWriter, d interface{}, url string) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
}

//Accepted : HTTP 202 Accepted
func Accepted(w http.ResponseWriter) {
	w.WriteHeader(http.StatusAccepted)
}

//NoContent : HTTP 204 response
func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte{})
}

//BadRequest : HTTP 400 response
func BadRequest(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(d)
}

//Unauthorized : HTTP 401
func Unauthorized(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(d)
}

//Forbidden : HTTP 403 response
func Forbidden(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	w.Write([]byte(`{"message" : "access denied"}`))
}

//NotFound : HTTP 404 response
func NotFound(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusNotFound)
	if d == nil {
		w.Write([]byte(`{"message" : "requested data is not found"}`))
	} else {
		json.NewEncoder(w).Encode(d)
	}
}

//PreconditionFailed : HTTP 412
func PreconditionFailed(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusPreconditionFailed)
	json.NewEncoder(w).Encode(d)
}
