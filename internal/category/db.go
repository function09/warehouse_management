package category

import (
	"database/sql"
	"fmt"
)

type PostGreSQLRepository struct {
	db *sql.DB
}

func NewPostGreSQLRepository(db *sql.DB) Repository {
	return &PostGreSQLRepository{db: db}
}

func (r *PostGreSQLRepository) GetCategoryByName(n string) (*Category, error) {
	sqlStatement := ("SELECT category_id, category_name FROM categories WHERE category_name = $1")

	row := r.db.QueryRow(sqlStatement, n)

	var category Category

	err := row.Scan(&category.CategoryID, &category.CategoryName)

	if err != nil {
		return nil, fmt.Errorf("Error querying category: %v", err)
	}

	return &category, nil
}

func (r *PostGreSQLRepository) GetCategoryByID(id int) (*Category, error) {
	sqlStatement := "SELECT category_id, category_name FROM categories WHERE category_id = $1"

	row := r.db.QueryRow(sqlStatement, id)

	var category Category

	err := row.Scan(&category.CategoryID, &category.CategoryName)

	if err != nil {
		return nil, fmt.Errorf("Error querying category: %v", err)
	}

	return &category, nil
}

func (r *PostGreSQLRepository) AddNewCategory(n string) (int64, error) {
	sqlStatement := "INSERT INTO categories (category_name) VALUES ($1) RETURNING category_id"

	var id int64
	err := r.db.QueryRow(sqlStatement, n).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Error inserting new row: %v", err)
	}

	return id, nil
}

func (r *PostGreSQLRepository) UpdateCategory(n string, id int) (int64, error) {
	sqlStatment := "UPDATE categories SET category_name=$1 WHERE category_id=$2 RETURNING category_id"

	var cat Category

	err := r.db.QueryRow(sqlStatment, n, id).Scan(&cat.CategoryName, &cat.CategoryID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("Category with ID %d not found: %w", id, err)
		} else {
			return 0, fmt.Errorf("Error querying category: %w", err)
		}
	}

	return int64(cat.CategoryID), nil
}

func (r *PostGreSQLRepository) DeleteCategory(id int) (int64, error) {
	sqlStatement := "DELETE FROM categories WHERE category_id = $1 RETURNING category_id"

	var catID int64

	err := r.db.QueryRow(sqlStatement, id).Scan(&catID)

	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("Category with ID %d not found: %w", catID, err)
		} else {
			return 0, fmt.Errorf("Error querying category: %w", err)
		}
	}

	return catID, nil
}
