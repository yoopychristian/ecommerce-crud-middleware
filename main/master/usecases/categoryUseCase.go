package usecases

import "ecommerce-crud-middleware/main/master/models"

type CategoryUseCase interface {
	GetAllCategories() ([]*models.Category, error)
	GetCategoryByID(Id string) (*models.Category, error)
	CreateCategories(category models.Category) error
	UpdateCategories(category models.Category) error
	DeleteCategories(Id string) error
}
