package mapperimpl

import (
	"github.com/feryadialoi/golang-ecommerce/entity"
	"github.com/feryadialoi/golang-ecommerce/model"
)

type ProductMapperImpl struct {
}

func (productMapper ProductMapperImpl) MapProductToGetProductResponse(product entity.Product, response *model.GetProductResponse) {
	response.Name = product.Name
	response.Price = product.Price
	response.Quantity = product.Quantity
}

func (productMapper ProductMapperImpl) MapProductToCreateProductResponse(product entity.Product, response *model.CreateProductResponse) {
	response.Name = product.Name
	response.Price = product.Price
	response.Quantity = product.Quantity
}

func (productMapper ProductMapperImpl) MapProductToUpdateProductResponse(product entity.Product, response *model.UpdateProductResponse) {
	response.Name = product.Name
	response.Price = product.Price
	response.Quantity = product.Quantity
}
