package repositories

import "ecommerce-crud-middleware/main/master/models"

type OrderRepository interface {
	GetAllOrders() ([]*models.Order, error)
	GetOrderByID(IdProduct string) (*models.Order, error)
	CreateOrders(order models.Order) error
	UpdateOrders(order models.Order) error
	DeleteOrders(IdProduct string) error
}
