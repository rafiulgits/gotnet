package main

import (
	"log"
	"net/http"

	"github.com/rafiulgits/gotnet"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response" : "ok"}`))
}

func main() {
	builder := gotnet.NewAppBuilder()

	app := builder.Build()

	app.MapHandlerFunc("/", helloHandler)

	app.BeforeRun(func() {
		log.Println("start background services")
	})

	app.Run(nil)
}
