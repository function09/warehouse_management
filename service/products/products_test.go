package service_test

import (
	"errors"
	"testing"

	domain "github.com/function09/warehouse_management/domain/products"
	service "github.com/function09/warehouse_management/service/products"
)

type MockRepository struct {
	products []*domain.Product
	err      error
}

func (m *MockRepository) GetProductByID(i int) (*domain.Product, error) {
	if m.err != nil {
		return nil, m.err
	}

	for _, p := range m.products {
		if p.ID == i {
			return p, nil
		}
	}

	return nil, errors.New("product not found")
}

func (m *MockRepository) GetProductByName(n string) (*domain.Product, error) {
	if m.err != nil {
		return nil, m.err
	}

	for _, p := range m.products {
		if p.Title == n {
			return p, nil
		}
	}

	return nil, errors.New("Not found")
}

func (m *MockRepository) GetAllProducts() ([]*domain.Product, error) {
	if m.err != nil {
		return nil, m.err
	}

	productsSlice := make([]*domain.Product, 0, len(m.products))

	for _, p := range m.products {
		productsSlice = append(productsSlice, p)
	}

	return productsSlice, nil
}

func TestGetProductByID(t *testing.T) {
	mockRepo := &MockRepository{
		products: []*domain.Product{
			{ID: 1, Title: "Laptop", Category: "Electronics", Stock: 99},
			{ID: 2, Title: "PlayStation", Category: "Electronics", Stock: 21},
		},
	}

	service := service.NewService(mockRepo)

	p, err := service.GetProductByID(2)

	if err != nil {
		t.Fatalf("Expected product with ID 2 but got %v", err)
	}

	if p == nil {
		t.Fatalf("Expected to get a product but got %v", p)
	}

	if p.Title != "PlayStation" {
		t.Fatalf("Expected to get Playstation but got %v", p.Title)
	}
}

func TestGetProductByName(t *testing.T) {
	mockRepo := &MockRepository{
		products: []*domain.Product{
			{ID: 1, Title: "Laptop", Category: "Electronics", Stock: 99},
			{ID: 2, Title: "PlayStation", Category: "Electronics", Stock: 21},
		},
	}

	service := service.NewService(mockRepo)

	p, err := service.GetProductByName("PlayStation")

	if err != nil {
		t.Fatalf("Expected product to be PlayStation but got %v", err)
	}

	if p.Title != "PlayStation" {
		t.Fatalf("Expected Title to be PlayStation but got %v", p.Title)
	}

}

func TestGetAllProducts(t *testing.T) {
	mockRepo := &MockRepository{
		products: []*domain.Product{
			{ID: 1, Title: "Laptop", Category: "Electronics", Stock: 99},
			{ID: 2, Title: "PlayStation", Category: "Electronics", Stock: 21},
		},
	}

	expectedProducts := []*domain.Product{
		{ID: 1, Title: "Laptop", Category: "Electronics", Stock: 99},
		{ID: 2, Title: "PlayStation", Category: "Electronics", Stock: 21},
	}

	service := service.NewService(mockRepo)

	p, err := service.GetAllProducts()

	if err != nil {
		t.Fatalf("Expected to get all products but got %v", err)
	}

	if len(mockRepo.products) != len(expectedProducts) {
		t.Errorf("Expected %v products but got %v", expectedProducts, mockRepo.products)
	}

	if len(p) == 0 {
		t.Errorf("Expected a list of products but %v", p)
	}
}
