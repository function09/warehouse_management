package product

import (
	"database/sql"
	"fmt"
	"strconv"
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

	err := row.Scan(&p.ID, &p.Category, &p.Name, &p.Stock)
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

		if err := rows.Scan(&p.ID, &p.Name, &p.Stock); err != nil {
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

		if err := rows.Scan(&p.ID, &p.Name, &p.Stock, &p.Category); err != nil {
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

		if err := rows.Scan(&p.ID, &p.Name, &p.Stock); err != nil {
			return nil, err
		}

		productList = append(productList, &p)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return productList, nil
}

func (r *PostgreSQLRepository) AddNewProduct(name string, stock int, category string) (int64, error) {
	// Must be modified to not allow duplicate entries
	sqlStatement := "INSERT INTO products (product_name, stock, category_id) VALUES ($1, $2, (SELECT category_id FROM categories WHERE category_name = $3)) RETURNING product_id"

	var id int64
	err := r.db.QueryRow(sqlStatement, name, stock, category).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("Error inserting products: %v", err)
	}

	return id, nil
}

func (r *PostgreSQLRepository) UpdateProduct(id int, name string, stock int, category string) (string, error) {
	sqlStatement := "UPDATE products SET product_name = $1, stock = $2, category_id = (SELECT category_id from categories WHERE category_name = $3) WHERE product_id = $4"

	result, err := r.db.Exec(sqlStatement, name, stock, category, id)

	if err != nil {
		return "", fmt.Errorf("Error updating products: %v", err)
	}

	rowsAffected, err := result.RowsAffected()

	if rowsAffected == 0 {
		return "", fmt.Errorf("No product found with ID %d to update", id)
	}

	return "Product " + strconv.Itoa(id) + " successfully updated", nil

}
