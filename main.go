package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("connected to database successfully")

	resp, err := http.Get("https://dummyjson.com/products/category-list")

	if err != nil {
		fmt.Println("Failed to fetch data:", err)
	}

	defer resp.Body.Close()

	var categories []string

	json.NewDecoder(resp.Body).Decode(&categories)

	for _, name := range categories {
		sqlStatement := `INSERT INTO categories(category_name) VALUES ($1)`
		_, err = db.Exec(sqlStatement, name)
		if err != nil {
			fmt.Println("Error inserting data", err)
		}
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		log.Fatal("APP_PORT not assigned")
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
