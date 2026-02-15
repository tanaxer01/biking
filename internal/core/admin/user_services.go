package admin

import "github.com/tanaxer01/biking/pkg/biking"

func (s *Service) GetAllUsers() ([]biking.User, error) {
	return s.userRepo.ListUsers()
}

func (s *Service) GetUser(ID int) (*biking.User, error) {
	return s.userRepo.UserByID(ID)
}

func (s *Service) UpdateUser(ID int, data biking.UserData) error {
	return s.userRepo.UpdateUserData(ID, data)
}
