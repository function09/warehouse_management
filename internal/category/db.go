package category

import "database/sql"

type PostGreSQLRepository struct {
	db *sql.DB
}

func NewPostGreSQLRepository(db *sql.DB) Repository {
	return &PostGreSQLRepository{db: db}
}

func (r *PostGreSQLRepository) GetCategoryByName(n string) (*Category, error) {
	return &Category{}, nil
}

func (r *PostGreSQLRepository) GetCategoryByID(id int) (*Category, error) {
	return &Category{}, nil
}
