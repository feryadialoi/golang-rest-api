package repository

import "github.com/feryadialoi/golang-ecommerce/entity"

type ProductRepository interface {
	FindAll() []entity.Product
	FindOneById(id int) entity.Product
	Save(product *entity.Product)
	DeleteById(id int)
}
