package usecases

import (
	"ecommerce-crud-middleware/main/master/models"
	"ecommerce-crud-middleware/main/master/repositories"
)

type ProductUsecaseImpl struct {
	productRepo repositories.ProductRepository
}

func (s ProductUsecaseImpl) GetProducts() ([]*models.Product, error) {
	product, err := s.productRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s ProductUsecaseImpl) GetProductByID(Id string) (*models.Product, error) {
	product, err := s.productRepo.GetProductByID(Id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s ProductUsecaseImpl) CreateProducts(product models.Product) error {
	err := s.productRepo.CreateProducts(product)
	if err != nil {
		return err
	}
	return nil
}

func (s ProductUsecaseImpl) UpdateProducts(product models.Product) error {
	err := s.productRepo.UpdateProducts(product)
	if err != nil {
		return err
	}
	return nil
}
func (s ProductUsecaseImpl) DeleteProducts(Id string) error {
	err := s.productRepo.DeleteProducts(Id)
	if err != nil {
		return err
	}
	return nil
}

func InitProductUseCase(productRepo repositories.ProductRepository) ProductUseCase {
	return &ProductUsecaseImpl{productRepo}
}
