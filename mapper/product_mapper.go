package mapper

import (
	"github.com/feryadialoi/golang-ecommerce/entity"
	"github.com/feryadialoi/golang-ecommerce/model"
)

type ProductMapper interface {
	MapProductToGetProductResponse(product entity.Product, response *model.GetProductResponse)
	MapProductToCreateProductResponse(product entity.Product, response *model.CreateProductResponse)
	MapProductToUpdateProductResponse(product entity.Product, response *model.UpdateProductResponse)
}
