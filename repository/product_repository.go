package repository

import "mountainio/entity"

type ProductRepository interface {
	Insert(product entity.Product)

	FindAll() (products []entity.Product)

	DeleteAll()
}
