package user

import (
	"time"

	"github.com/tanaxer01/biking/pkg/biking"
)

type Service struct {
	db     DB
	crypto Crypto
	auth   Auth
}

type DB interface {
	InsertUser(data biking.InsertUser) error
	UserByID(id int) (*biking.User, error)
	UserByEmail(email string) (*biking.User, error)
	UpdateUserData(ID int, data biking.UserData) error
}

type Crypto interface {
	HashPassword(pass string) (string, error)
}

type Auth interface {
	GenerateJwtToken(params map[string]any) (string, error)
	GetTokenClaim(tokenString string) (map[string]any, error)
}

func NewService(db DB, crypto Crypto, auth Auth) *Service {
	return &Service{db: db, crypto: crypto, auth: auth}
}

// * How do we diff between db error & already exists ??
// * Should we return a jwt ???
func (s *Service) Insert(data *biking.InsertUser) error {
	hashedPassword, err := s.crypto.HashPassword(data.Password)
	if err != nil {
		return err
	}

	data.Password = hashedPassword
	err = s.db.InsertUser(*data)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Login(data biking.LoginUser) (string, error) {
	user, err := s.db.UserByEmail(data.Email)
	if err != nil {
		return "", err
	}

	params := &map[string]any{
		"sub":       user.ID,
		"email":     user.Email,
		"firstName": user.FirstName,
		"lastName":  user.LastName,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token, err := s.auth.GenerateJwtToken(*params)
	if err != nil {
		return "", err
	}

	return token, nil
}

// * We have to get the info from the token
// * Check case with invalid sub
func (s *Service) GetProfile(tokenString string) (*biking.UserData, error) {
	claim, err := s.auth.GetTokenClaim(tokenString)
	if err != nil {
		return nil, err
	}

	id := int(claim["sub"].(float64))
	if id == 0 {
	}

	user, err := s.db.UserByID(id)
	if err != nil {
		return nil, err
	}

	return &user.UserData, nil
}

// * Same as getProfile with the id
// * Upgrade this method to allow partial updates
// * Updating the profile should invalidate the token shouldn't it ????
// * Also 30 days might be too long if profile data is going in the token
func (s *Service) UpdateProfile(tokenString string, input biking.UserData) error {
	claim, err := s.auth.GetTokenClaim(tokenString)
	if err != nil {
		return err
	}

	id := int(claim["sub"].(float64))
	if id == 0 {
	}

	return s.db.UpdateUserData(id, input)
}
