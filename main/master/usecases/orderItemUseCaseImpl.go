package usecases

import (
	"ecommerce-crud-middleware/main/master/models"
	"ecommerce-crud-middleware/main/master/repositories"
)

type OrderItemUsecaseImpl struct {
	orderItemRepo repositories.OrderItemRepository
}

func (s OrderItemUsecaseImpl) GetAllOrderItems() ([]*models.OrderItem, error) {
	orderItem, err := s.orderItemRepo.GetAllOrderItems()
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (s OrderItemUsecaseImpl) GetOrderItemByID(Id string) (*models.OrderItem, error) {
	orderItem, err := s.orderItemRepo.GetOrderItemByID(Id)
	if err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (s OrderItemUsecaseImpl) CreateOrderItems(orderItem models.OrderItem) error {
	err := s.orderItemRepo.CreateOrderItems(orderItem)
	if err != nil {
		return err
	}
	return nil
}

func (s OrderItemUsecaseImpl) UpdateOrderItems(orderItem models.OrderItem) error {
	err := s.orderItemRepo.UpdateOrderItems(orderItem)
	if err != nil {
		return err
	}
	return nil
}
func (s OrderItemUsecaseImpl) DeleteOrderItems(OrderId string) error {
	err := s.orderItemRepo.DeleteOrderItems(OrderId)
	if err != nil {
		return err
	}
	return nil
}

func InitOrderItemUseCase(orderItemRepo repositories.OrderItemRepository) OrderItemUseCase {
	return &OrderItemUsecaseImpl{orderItemRepo}
}
