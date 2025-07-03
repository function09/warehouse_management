package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/function09/warehouse_management/internal/product"
	_ "github.com/lib/pq"
)

func ConnectToDB() *sql.DB {

	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Println("connected to database successfully")

	return db
}

func AddCategories(db *sql.DB) error {
	resp, err := http.Get("https://dummyjson.com/products/category-list")

	if err != nil {
		return fmt.Errorf("Failed to fetch data: %w", err)
	}

	defer resp.Body.Close()

	var categories []string

	if err := json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return fmt.Errorf("Error decoding JSON: %w", err)
	}

	for _, name := range categories {
		sqlStatement := `INSERT INTO categories(category_name) VALUES ($1)`
		_, err = db.Exec(sqlStatement, name)
		if err != nil {
			fmt.Printf("Error inserting data %q: %v\n", name, err)
		}
	}

	return nil
}

type ProductsResponse struct {
	Products []Product `json:"products"`
}

type Product struct {
	Title    string `json:"title"`
	Category string `json:"category"`
	Stock    int    `json:"stock"`
}

func AddProducts(db *sql.DB) error {
	resp, err := http.Get("https://dummyjson.com/products?limit=0")

	if err != nil {
		return fmt.Errorf("Failed to fetch data: %w", err)
	}

	defer resp.Body.Close()

	var products ProductsResponse

	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return fmt.Errorf("Error decoding JSON: %w", err)
	}

	for _, p := range products.Products {
		sqlStatement := "INSERT INTO products (product_name, stock, category_id) VALUES($1, $2, (SELECT category_id FROM categories WHERE category_name = $3))"
		_, err := db.Exec(sqlStatement, p.Title, p.Stock, p.Category)
		if err != nil {
			return fmt.Errorf("Error inserting data: %w", err)
		}
	}
	return nil
}

func main() {
	db := ConnectToDB()
	defer db.Close()

	repo := product.NewPostgreSQLRepository(db)
	svc := product.NewService(repo)
	handler := product.NewProductHandler(svc)

	router := http.NewServeMux()

	router.HandleFunc("/products", handler.GetProductByID)
	router.HandleFunc("/names", handler.GetProductByName)
	router.HandleFunc("/all", handler.GetAllProducts)
	router.HandleFunc("/cat", handler.GetProductsByCategory)
	router.HandleFunc("/add", handler.AddNewProduct)
	router.HandleFunc("/update", handler.UpdateProduct)

	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("APP_PORT not assigned")
		port = "8080"
	}

	http.ListenAndServe(":"+port, router)
}
