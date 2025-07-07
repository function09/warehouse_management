package category

import (
	"fmt"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetCategoryByName(n string) (*Category, error) {
	if n == "" {
		return nil, fmt.Errorf("Category cannot be blank")
	}

	return s.repo.GetCategoryByName(n)
}

func (s *Service) GetCategoryByID(id int) (*Category, error) {
	return s.repo.GetCategoryByID(id)
}
