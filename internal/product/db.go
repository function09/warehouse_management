package product

import (
	"database/sql"
	"fmt"
)

type PostgreSQLRepository struct {
	db *sql.DB
}

func NewPostgreSQLRepository(db *sql.DB) Repository {
	return &PostgreSQLRepository{db: db}
}

func (r *PostgreSQLRepository) GetProductByID(id int) (*Product, error) {
	sqlStatement := "SELECT * FROM products WHERE product_id=$1"

	row := r.db.QueryRow(sqlStatement, id)

	var p Product

	err := row.Scan(&p.ID, &p.Category, &p.Title, &p.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with id %d not found", id)
		}
		return nil, err
	}
	return &p, nil
}

func (r *PostgreSQLRepository) GetProductByName(n string) (*Product, error) {
	sqlStatement := "SELECT * FROM products WHERE product_name=$1"

	row := r.db.QueryRow(sqlStatement, n)

	var p Product

	err := row.Scan(&p.ID, &p.Category, &p.Title, &p.Stock)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &p, nil
}

func (r *PostgreSQLRepository) GetAllProducts() ([]*Product, error) {
	sqlStatement := "select product_id, product_name, stock FROM products"

	rows, err := r.db.Query(sqlStatement, nil)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var productList []*Product
	for rows.Next() {
		var p Product

		if err := rows.Scan(&p.ID, &p.Title, &p.Stock); err != nil {
			return nil, err
		}

		productList = append(productList, &p)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productList, nil
}
