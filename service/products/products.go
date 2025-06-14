package productsservice

import (
	domain "github.com/function09/warehouse_management/domain/products"
)

type Service struct {
	repo domain.Repository
}

func NewService(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetProductByID(id int) (*domain.Product, error) {
	return s.repo.GetByID(id)
}
