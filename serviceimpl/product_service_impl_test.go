package serviceimpl

import (
	"fmt"
	"github.com/feryadialoi/golang-ecommerce/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

var productService = ProductServiceImpl{}

func TestProductServiceImpl_GetProducts(t *testing.T) {
	getProductResponses := productService.GetProducts()

	fmt.Println(getProductResponses)

	assert.NotEmpty(t, getProductResponses)
}

func TestProductServiceImpl_GetProduct(t *testing.T) {
	getProductResponse := productService.GetProduct(1)

	fmt.Println(getProductResponse)

	//assert.NotNil(t, getProductResponse)
	//assert.Equal(t, "Product 1", getProductResponse.Name)
}

func TestProductServiceImpl_CreateProduct(t *testing.T) {
	createProductRequest := model.CreateProductRequest{
		Name:     "Product 6",
		Price:    2000.00,
		Quantity: 20,
	}

	createProductResponse := productService.CreateProduct(createProductRequest)

	fmt.Println(createProductResponse)
	assert.Equal(t, "Product 6", createProductResponse.Name)
}
