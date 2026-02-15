package admin

import "github.com/tanaxer01/biking/pkg/biking"

type Service struct {
	bikeRepo BikeRepo
	userRepo UserRepo
}

type BikeRepo interface {
	InsertBike(biking.BikeData) error
	UpdateBikeData(ID int, data biking.BikeData) error
	ListBikes() ([]biking.Bike, error)
}

type UserRepo interface {
	ListUsers() ([]biking.User, error)
	UserByID(ID int) (*biking.User, error)
	UpdateUserData(ID int, data biking.UserData) error
}

func NewService(bikeRepo BikeRepo, userRepo UserRepo) *Service {
	return &Service{bikeRepo: bikeRepo, userRepo: userRepo}
}
