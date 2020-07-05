package repositories

import (
	"database/sql"
	"ecommerce-crud-middleware/main/master/models"
	"log"
)

type OrderItemRepoImpl struct {
	db *sql.DB
}

func (s OrderItemRepoImpl) GetAllOrderItems() ([]*models.OrderItem, error) {
	dataOrderItem := []*models.OrderItem{}
	query := "select * from purchase_order_item"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		orderItem := models.OrderItem{}
		var err = data.Scan(&orderItem.OrderId, &orderItem.Qty, &orderItem.ProductId, &orderItem.CreatedAt, &orderItem.UpdatedAt)
		if err != nil {
			return nil, err
		}
		dataOrderItem = append(dataOrderItem, &orderItem)
	}
	return dataOrderItem, nil
}

func (s OrderItemRepoImpl) GetOrderItemByID(OrderId string) (*models.OrderItem, error) {
	orderItem := new(models.OrderItem)
	query := "SELECT * FROM purchase_order_item WHERE order_id = ?"
	if err := s.db.QueryRow(query, OrderId).Scan(&orderItem.OrderId, &orderItem.Qty, &orderItem.ProductId, &orderItem.CreatedAt, &orderItem.UpdatedAt); err != nil {
		return nil, err
	}
	return orderItem, nil
}

func (s OrderItemRepoImpl) CreateOrderItems(orderItem models.OrderItem) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("INSERT INTO purchase_order_item VALUES (?,?,?,?,?)", orderItem.OrderId, orderItem.Qty, orderItem.ProductId, orderItem.CreatedAt, orderItem.UpdatedAt)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = stmt.LastInsertId()

	tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s OrderItemRepoImpl) UpdateOrderItems(orderItem models.OrderItem) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("UPDATE purchase_order_item SET qty=?, product_id=?, created_at=?, update_at=? WHERE order_id=?", orderItem.Qty, orderItem.ProductId, orderItem.CreatedAt, orderItem.UpdatedAt, orderItem.OrderId)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = stmt.LastInsertId()

	tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s OrderItemRepoImpl) DeleteOrderItems(OrderId string) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("DELETE FROM purchase_order_item WHERE order_id=?", OrderId)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	_, err = stmt.LastInsertId()

	tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func InitOrderItemRepoImpl(db *sql.DB) OrderItemRepository {
	return &OrderItemRepoImpl{db}
}
