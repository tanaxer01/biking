package admin

import "github.com/tanaxer01/biking/pkg/biking"

func (s *Service) InsertBike(data biking.BikeData) error {
	return s.bikeRepo.InsertBike(data)
}

func (s *Service) UpdateBike(ID int, data biking.BikeData) error {
	return s.bikeRepo.UpdateBikeData(ID, data)
}

func (s *Service) ListBikes() ([]biking.Bike, error) {
	return s.bikeRepo.ListBikes()
}
