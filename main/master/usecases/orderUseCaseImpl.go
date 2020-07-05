package usecases

import (
	"ecommerce-crud-middleware/main/master/models"
	"ecommerce-crud-middleware/main/master/repositories"
)

type OrderUseCaseImpl struct {
	orderRepo repositories.OrderRepository
}

func (s OrderUseCaseImpl) GetAllOrders() ([]*models.Order, error) {
	order, err := s.orderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s OrderUseCaseImpl) GetOrderByID(IdProduct string) (*models.Order, error) {
	order, err := s.orderRepo.GetOrderByID(IdProduct)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s OrderUseCaseImpl) CreateOrders(order models.Order) error {
	err := s.orderRepo.CreateOrders(order)
	if err != nil {
		return err
	}
	return nil
}

func (s OrderUseCaseImpl) UpdateOrders(order models.Order) error {
	err := s.orderRepo.UpdateOrders(order)
	if err != nil {
		return err
	}
	return nil
}
func (s OrderUseCaseImpl) DeleteOrders(IdProduct string) error {
	err := s.orderRepo.DeleteOrders(IdProduct)
	if err != nil {
		return err
	}
	return nil
}

func InitOrderUseCase(orderRepo repositories.OrderRepository) OrderUseCase {
	return &OrderUseCaseImpl{orderRepo}
}
