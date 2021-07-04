package main

import (
	"github.com/feryadialoi/golang-ecommerce/controller"
	"github.com/feryadialoi/golang-ecommerce/mapperimpl"
	"github.com/feryadialoi/golang-ecommerce/repositoryimpl"
	"github.com/feryadialoi/golang-ecommerce/serviceimpl"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()

	productController := controller.ProductController{
		ProductService: serviceimpl.ProductServiceImpl{
			ProductRepository: repositoryimpl.ProductRepositoryImpl{},
			ProductMapper:     mapperimpl.ProductMapperImpl{},
		},
	}

	router.GET("/products", productController.GetProducts)
	router.GET("/products/:productId", productController.GetProduct)
	router.POST("/products", productController.CreateProduct)
	router.PUT("/products/:productId", productController.UpdateProduct)
	router.DELETE("/products/:productId", productController.DeleteProduct)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
