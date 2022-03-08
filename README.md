# GotNET

A web framework designed for [NET](https://dotnet.microsoft.com/en-us/) friendly developers. Code base designed by following `.NET 6` minimal api structure.

GoNET framework is using following packages to provide the ultimate flavour 

* [uber-go/dig](https://github.com/uber-go/dig) dependency injection tool for golang
* [go-chi/chi](https://github.com/go-chi/chi) lightweight http service building tool for golang





### How to use

Install package in your project

`go get github.com/rafiulgits/gotnet`



Check [examples](https://github.com/rafiulgits/gotnet/tree/master/examples) 



**Simple http server**

```go
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"response" : "ok"}`))
}


func main() {
	builder := gotnet.NewAppBuilder()

	app := builder.Build()

	app.MapHandlerFunc("/", helloHandler)

	app.Run()
}

```





