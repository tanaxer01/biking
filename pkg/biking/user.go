package biking

type UserData struct {
	Email     string `json:"email" validate:"email,required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
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
	ID             int    `json:"id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	HashedPassword string
	UserData
}
