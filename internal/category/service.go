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
	if id < 0 {
		return nil, fmt.Errorf("Category ID cannot be less than 0")
	}

	return s.repo.GetCategoryByID(id)
}

func (s *Service) AddNewCategory(n string) (int64, error) {
	if n == "" {
		return 0, fmt.Errorf("category name cannot be blank")
	}

	return s.repo.AddNewCategory(n)
}

func (s *Service) UpdateCategory(n string, id int) (int64, error) {
	if n == "" {
		return 0, fmt.Errorf("category name cannot be blank")
	}

	return s.repo.UpdateCategory(n, id)
}

// func (s *Service) DeleteCategory(id int) (int64, error) {
// 	if id < 0 {
// 		return 0, fmt.Errorf("category ID cannot be less than 0")
// 	}
//
// 	return s.repo.DeleteCategory(id)
// }
