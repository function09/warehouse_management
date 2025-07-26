package shipment

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetShipments(l int, o int) ([]*Shipment, error) {

	if l < 0 || o < 0 {
		return nil, fmt.Errorf("Error: Limit and offset cannot be less than 0")
	}

	return s.repo.GetShipments(l, o)
}
