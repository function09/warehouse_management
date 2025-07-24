package shipment

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetShipments() ([]*Shipment, error) {
	return s.repo.GetShipments()
}
