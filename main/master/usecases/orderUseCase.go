package usecases

import "ecommerce-crud-middleware/main/master/models"

type OrderUseCase interface {
	GetAllOrders() ([]*models.Order, error)
	GetOrderByID(Id string) (*models.Order, error)
	CreateOrders(order models.Order) error
	UpdateOrders(order models.Order) error
	DeleteOrders(Id string) error
}
