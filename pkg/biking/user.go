package biking

import "time"

type UserData struct {
	Email     string `json:"email" validate:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type InsertUser struct {
	Password string
	UserData
}

type LoginUser struct {
	Email    string
	Password string
}

type User struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserData
}
