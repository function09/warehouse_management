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
	return &Category{}, nil
}
