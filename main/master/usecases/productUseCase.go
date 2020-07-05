package usecases

import "ecommerce-crud-middleware/main/master/models"

type ProductUseCase interface {
	GetProducts() ([]*models.Product, error)
	GetProductByID(Id string) (*models.Product, error)
	CreateProducts(student models.Product) error
	UpdateProducts(student models.Product) error
	DeleteProducts(Id string) error
}
