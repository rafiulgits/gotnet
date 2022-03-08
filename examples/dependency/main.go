package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rafiulgits/gotnet"
)

func main() {
	builder := gotnet.NewAppBuilder()

	app := builder.Build()

	app.Service.AddSingleton(NewProductService)

	app.RegisterHandler(NewProductHandler, func(router *chi.Mux, container *gotnet.Container) {
		container.Invoke(func(handler IProductHandler) {
			router.Route("/products", handler.Handle)
		})
	})

	app.Run()
}

// Handler Layer
type IProductHandler interface {
	gotnet.IHandler
}

type ProductHandler struct {
	productService IProductService
}

func NewProductHandler(productService IProductService) IProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (handler *ProductHandler) Handle(router chi.Router) {
	router.Get("/", handler.getProductHandler)
}

func (handler *ProductHandler) getProductHandler(w http.ResponseWriter, r *http.Request) {
	product := handler.productService.GetProduct()
	gotnet.Ok(w, product)
}

// Service Layer
type IProductService interface {
	GetProduct() *ProductModel
}

type ProductService struct {
}

func NewProductService() IProductService {
	return &ProductService{}
}

func (service *ProductService) GetProduct() *ProductModel {
	return &ProductModel{ID: 1, Name: "Mango"}
}

// Data Models
type ProductModel struct {
	ID   int
	Name string
}
