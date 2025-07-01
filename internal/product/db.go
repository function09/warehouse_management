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

func (r *PostgreSQLRepository) GetProductByName(n string) ([]*Product, error) {
	sqlStatement := "SELECT product_id, product_name, stock FROM products WHERE product_name ILIKE $1"

	query := "%" + n + "%"

	rows, err := r.db.Query(sqlStatement, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var productList []*Product

	for rows.Next() {
		p := &Product{}

		if err := rows.Scan(&p.ID, &p.Title, &p.Stock); err != nil {
			return nil, err
		}

		productList = append(productList, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return productList, nil
}

func (r *PostgreSQLRepository) GetProductsByCategory(c string) ([]*Product, error) {
	sqlStatement := "SELECT p.product_id, p.product_name, p.stock, c.category_name FROM products as p INNER JOIN categories as c ON p.category_id = c.category_id WHERE c.category_name = $1;"

	rows, err := r.db.Query(sqlStatement, c)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []*Product

	for rows.Next() {
		p := &Product{}

		if err := rows.Scan(&p.ID, &p.Title, &p.Stock, &p.Category); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *PostgreSQLRepository) GetAllProducts(limit int, offset int) ([]*Product, error) {
	sqlStatement := "SELECT product_id, product_name, stock FROM products ORDER BY product_id LIMIT $1 OFFSET $2"

	rows, err := r.db.Query(sqlStatement, limit, offset)

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
