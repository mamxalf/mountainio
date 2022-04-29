package repository

import (
	"mountainio/domain/entity"
)

type ProductRepository interface {
	Insert(product entity.Product)

	FindAll() (products []entity.Product)

	DeleteAll()
}
