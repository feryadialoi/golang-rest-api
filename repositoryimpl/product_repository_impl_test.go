package repositoryimpl

import (
	"fmt"
	"github.com/feryadialoi/golang-ecommerce/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

var productRepository = ProductRepositoryImpl{}

func TestProductRepositoryImpl_FindAll(t *testing.T) {
	products := productRepository.FindAll()
	fmt.Println(products)

	assert.NotEmpty(t, products)
}
func TestProductRepositoryImpl_FindOneById(t *testing.T) {
	product1 := productRepository.FindOneById(1)
	fmt.Println(product1)

	assert.NotNil(t, product1)
	assert.Equal(t, 1, product1.ID)
	assert.Equal(t, "Product 1", product1.Name)

	product2 := productRepository.FindOneById(2)
	fmt.Println(product2)

	assert.NotNil(t, product2)
	assert.Equal(t, 2, product2.ID)
	assert.Equal(t, "Product 2", product2.Name)
}

func TestProductRepositoryImpl_Save(t *testing.T) {
	product := entity.Product{
		Name:     "Product test",
		Price:    150000,
		Quantity: 20,
	}
	fmt.Println(product)
	assert.Equal(t, 0, product.ID)
	productRepository.Save(&product)

	fmt.Println(product)
	assert.Equal(t, 6, product.ID)
}

func TestProductRepositoryImpl_DeleteById(t *testing.T) {
	productId := 1
	products := productRepository.FindAll()
	fmt.Println(products)

	productRepository.DeleteById(productId)
	products = productRepository.FindAll()
	fmt.Println(products)
	assert.Equal(t, 4, len(products))
}
