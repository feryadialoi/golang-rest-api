package service

import "github.com/feryadialoi/golang-ecommerce/model"

type ProductService interface {
	GetProducts() []model.GetProductResponse
	GetProduct(id int) model.GetProductResponse
	CreateProduct(request model.CreateProductRequest) model.CreateProductResponse
	UpdateProduct(id int, request model.UpdateProductRequest) model.UpdateProductResponse
	DeleteProduct(id int)
}
