package main

import (
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

	app.MapHandler("/", helloHandler)

	app.Run()
}
