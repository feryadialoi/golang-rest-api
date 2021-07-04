package mapperimpl

import (
	"fmt"
	"github.com/feryadialoi/golang-ecommerce/entity"
	"github.com/feryadialoi/golang-ecommerce/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var productMapper = ProductMapperImpl{}

func TestProductMapperImpl_MapProductToGetProductResponse(t *testing.T) {
	product := entity.Product{
		ID:       1,
		Name:     "Product 1",
		Price:    10000,
		Quantity: 10,
	}

	getProductResponse := model.GetProductResponse{}
	fmt.Println(getProductResponse)

	productMapper.MapProductToGetProductResponse(product, &getProductResponse)

	fmt.Println(getProductResponse)

	assert.Equal(t, "Product 1", getProductResponse.Name)
	assert.Equal(t, float64(10000), getProductResponse.Price)
	assert.Equal(t, 10, getProductResponse.Quantity)
}

func TestProductMapperImpl_MapProductToCreateProductResponse(t *testing.T) {
	product := entity.Product{
		ID:       1,
		Name:     "Product 1",
		Price:    10000,
		Quantity: 10,
	}

	createProductResponse := model.CreateProductResponse{}
	fmt.Println(createProductResponse)

	productMapper.MapProductToCreateProductResponse(product, &createProductResponse)

	fmt.Println(createProductResponse)
	assert.Equal(t, "Product 1", createProductResponse.Name)
	assert.Equal(t, float64(10000), createProductResponse.Price)
	assert.Equal(t, 10, createProductResponse.Quantity)

}

func TestProductMapperImpl_MapProductToUpdateProductResponse(t *testing.T) {
	product := entity.Product{
		ID:       1,
		Name:     "Product 1",
		Price:    10000,
		Quantity: 10,
	}

	updateProductResponse := model.UpdateProductResponse{}
	fmt.Println(updateProductResponse)

	productMapper.MapProductToUpdateProductResponse(product, &updateProductResponse)

	fmt.Println(updateProductResponse)
	assert.Equal(t, "Product 1", updateProductResponse.Name)
	assert.Equal(t, float64(10000), updateProductResponse.Price)
	assert.Equal(t, 10, updateProductResponse.Quantity)

}
