package repositories

import (
	"database/sql"
	"ecommerce-crud-middleware/main/master/models"
	"log"
)

type CategoryRepoImpl struct {
	db *sql.DB
}

func (s CategoryRepoImpl) GetAllCategories() ([]*models.Category, error) {
	dataCategory := []*models.Category{}
	query := "select * from category"
	data, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	for data.Next() {
		category := models.Category{}
		var err = data.Scan(&category.Id, &category.CategoryName, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		dataCategory = append(dataCategory, &category)
	}
	return dataCategory, nil
}

func (s CategoryRepoImpl) GetCategoryByID(Id string) (*models.Category, error) {
	category := new(models.Category)
	query := "SELECT * FROM category WHERE category_id = ?"
	if err := s.db.QueryRow(query, Id).Scan(&category.Id, &category.CategoryName, &category.CreatedAt, &category.UpdatedAt); err != nil {
		return nil, err
	}
	return category, nil
}

func (s CategoryRepoImpl) CreateCategories(category models.Category) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("INSERT INTO category VALUES (?,?,?,?)", category.Id, category.CategoryName, category.CreatedAt, category.UpdatedAt)
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

func (s CategoryRepoImpl) UpdateCategories(category models.Category) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("UPDATE category SET  category_name=?, created_at=?, update_at=? WHERE category_id=?", category.CategoryName, category.CreatedAt, category.UpdatedAt, category.Id)
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

func (s CategoryRepoImpl) DeleteCategories(Id string) error {
	tx, err := s.db.Begin()
	stmt, err := tx.Exec("DELETE FROM category WHERE category_id=?", Id)
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

func InitCategoryRepoImpl(db *sql.DB) CategoryRepository {
	return &CategoryRepoImpl{db}
}
