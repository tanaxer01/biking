package sqlite

import (
	"database/sql"
	"time"

	"github.com/tanaxer01/biking/pkg/biking"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// * Add transacctions to avoid issues
func (r *UserRepository) InsertUser(data biking.InsertUser) error {
	_, err := r.db.Exec(
		`INSERT INTO users (email, first_name, last_name, hashed_password) VALUES (?, ?, ?, ?)`,
		data.Email,
		data.FirstName,
		data.LastName,
		data.Password,
	)

	return err
}

func (r *UserRepository) UpdateUserData(ID int, data biking.UserData) error {
	_, err := r.db.Exec(
		`UPDATE users SET email = ?, first_name = ?, last_name = ?, updated_at = ? WHERE id = ?`,
		data.Email,
		data.FirstName,
		data.LastName,
		time.Now(),
		ID,
	)

	return err
}

func (r *UserRepository) ListUsers() ([]biking.User, error) {
	var users []biking.User

	rows, err := r.db.Query("SELECT id, email, first_name, last_name, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user biking.User
		err = rows.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) UserByID(id int) (*biking.User, error) {
	var user biking.User
	err := r.db.
		QueryRow("SELECT id, email, first_name, last_name, created_at, updated_at FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UserByEmail(email string) (*biking.User, error) {
	var user biking.User
	err := r.db.
		QueryRow("SELECT id, email, first_name, last_name, created_at, updated_at FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
