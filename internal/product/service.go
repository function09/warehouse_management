package product

import (
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProductByID(id int) (*Product, error) {
	if id <= 0 {
		return nil, errors.New("Error: ID cannot be 0 or negative.")
	}

	product, err := s.repo.GetProductByID(id)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *Service) GetProductByName(n string) (*Product, error) {
	return s.repo.GetProductByName(n)
}

func (s *Service) GetAllProducts() ([]*Product, error) {
	return s.repo.GetAllProducts()
}
