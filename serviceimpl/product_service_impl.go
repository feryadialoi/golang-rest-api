package serviceimpl

import (
	"github.com/feryadialoi/golang-ecommerce/entity"
	"github.com/feryadialoi/golang-ecommerce/mapper"
	"github.com/feryadialoi/golang-ecommerce/model"
	"github.com/feryadialoi/golang-ecommerce/repository"
)

type ProductServiceImpl struct {
	repository.ProductRepository
	mapper.ProductMapper
}

func (productService ProductServiceImpl) GetProducts() []model.GetProductResponse {
	products := productService.ProductRepository.FindAll()

	var getProductResponses []model.GetProductResponse
	for _, product := range products {
		getProductResponse := model.GetProductResponse{}
		productService.ProductMapper.MapProductToGetProductResponse(product, &getProductResponse)
		getProductResponses = append(getProductResponses, getProductResponse)
	}

	return getProductResponses
}

func (productService ProductServiceImpl) GetProduct(id int) model.GetProductResponse {
	product := productService.ProductRepository.FindOneById(id)

	var getProductResponse model.GetProductResponse

	productService.ProductMapper.MapProductToGetProductResponse(product, &getProductResponse)

	return getProductResponse
}

func (productService ProductServiceImpl) CreateProduct(request model.CreateProductRequest) model.CreateProductResponse {
	product := entity.Product{}
	product.Name = request.Name
	product.Price = request.Price
	product.Quantity = request.Quantity

	productService.ProductRepository.Save(&product)

	createProductResponse := model.CreateProductResponse{}
	productService.ProductMapper.MapProductToCreateProductResponse(product, &createProductResponse)

	return createProductResponse
}

func (productService ProductServiceImpl) UpdateProduct(id int, request model.UpdateProductRequest) model.UpdateProductResponse {
	product := productService.ProductRepository.FindOneById(id)
	product.Name = request.Name
	product.Price = request.Price
	product.Quantity = request.Quantity

	productService.ProductRepository.Save(&product)

	updateProductResponse := model.UpdateProductResponse{}
	productService.ProductMapper.MapProductToUpdateProductResponse(product, &updateProductResponse)

	return updateProductResponse
}

func (productService ProductServiceImpl) DeleteProduct(id int) {
	productService.ProductRepository.DeleteById(id)
}
