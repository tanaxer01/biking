package rental

type Service struct {
	db DB
}

type DB interface {
}

func NewService(db DB) *Service {
	return &Service{db: db}
}

func (s *Service) StartRental() error {
	return nil
}

func (s *Service) EndRental() error {
	return nil
}

func (s *Service) GetHistory() error {
	return nil
}

func (s *Service) GetAll() error {
	return nil
}

func (s *Service) GetDetails() error {
	return nil
}

func (s *Service) UpdateDetails() error {
	return nil
}
