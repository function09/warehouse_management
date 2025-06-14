package postgres

import (
	"database/sql"

	"github.com/function09/warehouse_management/domain/products"
)

type PostGreSQLRepository struct {
	db *sql.DB
}

func newPostGreSQLRepository(db *sql.DB) *PostGreSQLRepository {
	return &PostGreSQLRepository{db: db}
}

// func (r *PostGreSQLRepository) GetByID(id int) (*products.Product, error) {
// }

