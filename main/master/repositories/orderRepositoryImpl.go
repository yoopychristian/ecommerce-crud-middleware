package repositories

import (
	"database/sql"
	"ecommerce-crud-middleware/main/master/models"
	"log"
)

type OrderRepoImpl struct {
	db *sql.DB
}

func (s OrderRepoImpl) GetAllOrders() ([]*models.Order, error) {
	dataOrder := []*models.Order{}
	query := "select * from purchase_order"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		order := models.Order{}
		var err = data.Scan(&order.IdProduct, &order.OrderDate, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			return nil, err
		}
		dataOrder = append(dataOrder, &order)
	}
	return dataOrder, nil
}

func (s OrderRepoImpl) GetOrderByID(IdProduct string) (*models.Order, error) {
	order := new(models.Order)
	query := "SELECT * FROM purchase_order WHERE order_id = ?"
	if err := s.db.QueryRow(query, IdProduct).Scan(&order.IdProduct, &order.OrderDate, &order.CreatedAt, &order.UpdatedAt); err != nil {
		return nil, err
	}
	return order, nil
}

func (s OrderRepoImpl) CreateOrders(order models.Order) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("INSERT INTO purchase_order VALUES (?,?,?,?)", order.IdProduct, order.OrderDate, order.CreatedAt, order.UpdatedAt)
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

func (s OrderRepoImpl) UpdateOrders(order models.Order) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("UPDATE purchase_order SET  order_Date=?, created_at=?, update_at=? WHERE id_product=?", order.OrderDate, order.CreatedAt, order.UpdatedAt, order.IdProduct)
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

func (s OrderRepoImpl) DeleteOrders(IdProduct string) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("DELETE FROM purchase_order WHERE order_id=?", IdProduct)
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

func InitOrderRepoImpl(db *sql.DB) OrderRepository {
	return &OrderRepoImpl{db}
}
