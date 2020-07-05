package repositories

import "ecommerce-crud-middleware/main/master/models"

type ProductRepository interface {
	GetAllProduct() ([]*models.Product, error)
	GetProductByID(Id string) (*models.Product, error)
	CreateProducts(product models.Product) error
	UpdateProducts(product models.Product) error
	DeleteProducts(Id string) error
}
