package service

import "mountainio/model"

type ProductService interface {
	Create(request model.CreateProductRequest) (response model.CreateProductResponse)
	List() (responses []model.GetProductResponse)
}
