package controller

import (
	"encoding/json"
	"fmt"
	"github.com/feryadialoi/golang-ecommerce/model"
	"github.com/feryadialoi/golang-ecommerce/service"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"
)

type ProductController struct {
	service.ProductService
}

func (productController ProductController) GetProducts(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	getProductResponses := productController.ProductService.GetProducts()

	jsonGetProductResponses, _ := json.Marshal(getProductResponses)

	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonGetProductResponses)

}

func (productController ProductController) GetProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId, _ := strconv.Atoi(params.ByName("productId"))
	getProductResponse := productController.ProductService.GetProduct(productId)

	jsonGetProductResponse, _ := json.Marshal(getProductResponse)

	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonGetProductResponse)
}

func (productController ProductController) CreateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	body, _ := io.ReadAll(request.Body)

	createProductRequest := model.CreateProductRequest{}

	json.Unmarshal(body, &createProductRequest)

	fmt.Println(createProductRequest)

	createProductResponse := productController.ProductService.CreateProduct(createProductRequest)

	jsonCreateProductResponse, _ := json.Marshal(createProductResponse)
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonCreateProductResponse)
}

func (productController ProductController) UpdateProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	body, _ := io.ReadAll(request.Body)

	productId, _ := strconv.Atoi(params.ByName("productId"))
	updateProductRequest := model.UpdateProductRequest{}

	json.Unmarshal(body, &updateProductRequest)

	fmt.Println(updateProductRequest)

	updateProductResponse := productController.ProductService.UpdateProduct(productId, updateProductRequest)

	jsonCreateProductResponse, _ := json.Marshal(updateProductResponse)
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(jsonCreateProductResponse)
}

func (productController ProductController) DeleteProduct(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId, _ := strconv.Atoi(params.ByName("productId"))
	productController.ProductService.DeleteProduct(productId)

	writer.Header().Add("Content-Type", "application/json")
	jsonData, _ := json.Marshal(productController.ProductService.GetProducts())
	writer.Write(jsonData)
}
