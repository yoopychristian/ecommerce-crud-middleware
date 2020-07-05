package repositories

import "ecommerce-crud-middleware/main/master/models"

type CategoryRepository interface {
	GetAllCategories() ([]*models.Category, error)
	GetCategoryByID(Id string) (*models.Category, error)
	CreateCategories(category models.Category) error
	UpdateCategories(category models.Category) error
	DeleteCategories(Id string) error
}
