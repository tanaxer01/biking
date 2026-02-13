package bike

import "github.com/tanaxer01/biking/pkg/biking"

type Service struct {
	db DB
}

type DB interface {
}

func NewService(db DB) *Service {
	return &Service{db: db}
}

func (s *Service) GetAvailable() ([]biking.BikeData, error) {
	return nil, nil
}
