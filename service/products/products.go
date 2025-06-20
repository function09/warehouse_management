package service

import (
	"errors"
	domain "github.com/function09/warehouse_management/domain/products"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProductByID(id int) (*domain.Product, error) {
	if id <= 0 {
		return nil, errors.New("Error: ID cannot be 0 or negative.")
	}

	product, err := s.repo.GetProductByID(id)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) GetProductByName(n string) (*domain.Product, error) {
	return s.repo.GetProductByName(n)
}

func (s *Service) GetAllProducts() ([]*domain.Product, error) {
	return s.repo.GetAllProducts()
}
