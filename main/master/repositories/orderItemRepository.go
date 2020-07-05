package repositories

import "ecommerce-crud-middleware/main/master/models"

type OrderItemRepository interface {
	GetAllOrderItems() ([]*models.OrderItem, error)
	GetOrderItemByID(OrderId string) (*models.OrderItem, error)
	CreateOrderItems(orderItem models.OrderItem) error
	UpdateOrderItems(orderItem models.OrderItem) error
	DeleteOrderItems(OrderId string) error
}
