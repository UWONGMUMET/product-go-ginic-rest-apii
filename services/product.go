package services

import "gin_gonic_products_rest_api/model"

type ProductServiceInterface interface {
	CreateProduct(model.PostProduct) error
	GetAllProducts() ([]model.Product, error)
	GetOneProduct(uint) (*model.Product, error)
	UpdateProduct(uint, model.PostProduct) error
	DeleteProduct(uint) error
}
