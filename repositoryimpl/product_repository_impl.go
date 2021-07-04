package repositoryimpl

import "github.com/feryadialoi/golang-ecommerce/entity"

var sequence = 5
var products []entity.Product = []entity.Product{
	{
		ID:       1,
		Name:     "Product 1",
		Price:    10000,
		Quantity: 10,
	}, {
		ID:       2,
		Name:     "Product 2",
		Price:    10000,
		Quantity: 10,
	}, {
		ID:       3,
		Name:     "Product 3",
		Price:    10000,
		Quantity: 10,
	}, {
		ID:       4,
		Name:     "Product 4",
		Price:    10000,
		Quantity: 10,
	}, {
		ID:       5,
		Name:     "Product 5",
		Price:    10000,
		Quantity: 10,
	},
}

type ProductRepositoryImpl struct {
}

func (productRepo ProductRepositoryImpl) FindAll() []entity.Product {
	return products
}

func (productRepo ProductRepositoryImpl) FindOneById(id int) entity.Product {
	product := entity.Product{}

	for _, v := range products {
		if v.ID == id {
			product = v
		}
	}

	return product
}

func (productRepo ProductRepositoryImpl) Save(product *entity.Product) {
	sequence = sequence + 1
	product.ID = sequence
	products = append(products, *product)
}

func (productRepo ProductRepositoryImpl) DeleteById(id int) {
	var index int
	for k, v := range products {
		if v.ID == id {
			index = k
		}
	}

	products = append(products[:index], products[index+1:]...)
}
