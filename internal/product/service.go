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

func (s *Service) GetProductByName(n string) ([]*Product, error) {
	product, err := s.repo.GetProductByName(n)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *Service) GetProductsByCategory(c string) ([]*Product, error) {
	products, err := s.repo.GetProductsByCategory(c)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *Service) GetAllProducts(limit int, offset int) ([]*Product, error) {
	if limit < 0 || offset < 0 {
		return nil, errors.New("Error: Limit and offset cannot be less than 0")
	}

	products, err := s.repo.GetAllProducts(limit, offset)

	if err != nil {
		return nil, err
	}

	return products, nil
}
