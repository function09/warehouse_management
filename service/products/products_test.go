package service_test

import (
	"errors"
	"testing"

	domain "github.com/function09/warehouse_management/domain/products"
	productsService "github.com/function09/warehouse_management/service/products"
)

type MockRepository struct {
	products map[int]*domain.Product
	err      error
}

func (m *MockRepository) GetProductByID(i int) (*domain.Product, error) {
	if m.err != nil {
		return nil, m.err
	}
	p, ok := m.products[i]

	if !ok {
		return nil, errors.New("not found")
	}

	return p, nil
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

func TestGetProductByIDSuccess(t *testing.T) {
	mockRepo := &MockRepository{products: map[int]*domain.Product{1: {ID: 1, Title: "Laptop", Category: "Electronics", Stock: 99}}}

	service := productsService.NewService(mockRepo)

	product, err := service.GetProductByID(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if product.ID != 1 || product.Title != "Laptop" {
		t.Errorf("receieved unexpected product: %+v", product)
	}
}

func TestGetProductByIDFail(t *testing.T) {
	mockRepo := &MockRepository{products: map[int]*domain.Product{}}

	service := productsService.NewService(mockRepo)

	_, err := service.GetProductByID(99)

	if err == nil {
		t.Fatalf("expected to get a product, got nil")
	}

}

func TestGetProductByID_RepositoryError(t *testing.T) {
	mockRepo := &MockRepository{err: errors.New("db connection failed")}

	service := productsService.NewService(mockRepo)

	_, err := service.GetProductByID(99)

	if err == nil || err.Error() != "db connection failed" {
		t.Fatalf("expected db connection failed error, got %v", err)
	}
}

func TestGetProductByNameSuccess(t *testing.T) {
	mockRepo := &MockRepository{products: map[int]*domain.Product{1: {ID: 1, Title: "Playstation", Category: "Electronics", Stock: 99}, 2: {ID: 1, Title: "Laptop", Category: "Electronics", Stock: 21}}}

	service := productsService.NewService(mockRepo)

	product, err := service.GetProductByName("Laptop")

	if err != nil {
		t.Errorf("Expected product %v got something else", err)
	}

	if product.Title != "Laptop" {
		t.Errorf("expected Laptop got %v", product.Title)
	}
}

func TestGetProductByNameFail(t *testing.T) {
	mockRepo := &MockRepository{products: map[int]*domain.Product{}}

	service := productsService.NewService(mockRepo)

	_, err := service.GetProductByName("Laptop")

	if err == nil {
		t.Errorf("Expected to get a product, got nil")
	}
}
