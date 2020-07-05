package usecases

import (
	"ecommerce-crud-middleware/main/master/models"
	"ecommerce-crud-middleware/main/master/repositories"
)

type CategoryUsecaseImpl struct {
	categoryRepo repositories.CategoryRepository
}

func (s CategoryUsecaseImpl) GetAllCategories() ([]*models.Category, error) {
	category, err := s.categoryRepo.GetAllCategories()
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryUsecaseImpl) GetCategoryByID(Id string) (*models.Category, error) {
	category, err := s.categoryRepo.GetCategoryByID(Id)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryUsecaseImpl) CreateCategories(category models.Category) error {
	err := s.categoryRepo.CreateCategories(category)
	if err != nil {
		return err
	}
	return nil
}

func (s CategoryUsecaseImpl) UpdateCategories(category models.Category) error {
	err := s.categoryRepo.UpdateCategories(category)
	if err != nil {
		return err
	}
	return nil
}
func (s CategoryUsecaseImpl) DeleteCategories(Id string) error {
	err := s.categoryRepo.DeleteCategories(Id)
	if err != nil {
		return err
	}
	return nil
}

func InitCategoryUseCase(categoryRepo repositories.CategoryRepository) CategoryUseCase {
	return &CategoryUsecaseImpl{categoryRepo}
}
