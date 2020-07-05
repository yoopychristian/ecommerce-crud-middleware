package usecases

import "ecommerce-crud-middleware/main/master/models"

type OrderItemUseCase interface {
	GetAllOrderItems() ([]*models.OrderItem, error)
	GetOrderItemByID(Id string) (*models.OrderItem, error)
	CreateOrderItems(orderItem models.OrderItem) error
	UpdateOrderItems(orderItem models.OrderItem) error
	DeleteOrderItems(Id string) error
}
