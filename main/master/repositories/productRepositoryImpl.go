package repositories

import (
	"database/sql"
	"ecommerce-crud-middleware/main/master/models"
	"log"
)

type ProductRepoImpl struct {
	db *sql.DB
}

func (s ProductRepoImpl) GetAllProduct() ([]*models.Product, error) {
	dataProduct := []*models.Product{}
	query := "select * from product"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		product := models.Product{}
		var err = data.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.CategoryID, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		dataProduct = append(dataProduct, &product)
	}
	return dataProduct, nil
}

func (s ProductRepoImpl) GetProductByID(Id string) (*models.Product, error) {
	product := new(models.Product)
	query := "SELECT * FROM product WHERE id = ?"
	if err := s.db.QueryRow(query, Id).Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.CategoryID, &product.Price, &product.CreatedAt, &product.UpdatedAt); err != nil {
		return nil, err
	}
	return product, nil
}

func (s ProductRepoImpl) CreateProducts(product models.Product) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("INSERT INTO product VALUES (?,?,?,?,?,?,?)", product.Id, product.ProductCode, product.ProductName, product.CategoryID, product.Price, product.CreatedAt, product.UpdatedAt)
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

func (s ProductRepoImpl) UpdateProducts(product models.Product) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("UPDATE product SET product_code=?, product_name=?, category_id=?, created_at=?, update_at=? WHERE id=?", product.ProductCode, product.ProductName, product.CategoryID, product.Price, product.CreatedAt, product.UpdatedAt, product.Id)
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

func (s ProductRepoImpl) DeleteProducts(Id string) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("DELETE FROM product WHERE id=?", Id)
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

func InitProductRepoImpl(db *sql.DB) ProductRepository {
	return &ProductRepoImpl{db}
}
