package sqlite

import (
	"database/sql"
	"time"

	"github.com/tanaxer01/biking/pkg/biking"
)

type BikeRepository struct {
	db *sql.DB
}

func NewBikeRepository(db *sql.DB) *BikeRepository {
	return &BikeRepository{db: db}
}

func (r *BikeRepository) InsertBike(data biking.BikeData) error {
	_, err := r.db.Exec(
		`INSERT INTO bikes (latitude, longitude) VALUES (?, ?)`,
		data.Latitude,
		data.Longitude,
	)

	return err
}

func (r *BikeRepository) UpdateBikeData(ID int, data biking.BikeData) error {
	_, err := r.db.Exec(
		`UPDATE bikes SET is_available = ?, latitude = ?, longitude = ?, updated_at = ? WHERE id = ?`,
		data.Available,
		data.Latitude,
		data.Longitude,
		time.Now(),
		ID,
	)

	return err
}

func (r *BikeRepository) ListBikes() ([]biking.Bike, error) {
	var bikes []biking.Bike

	rows, err := r.db.Query("SELECT id, is_available, latitude, longitude, created_at, updated_at FROM bikes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bike biking.Bike
		err = rows.Scan(&bike.ID, &bike.Available, &bike.Latitude, &bike.Longitude, &bike.CreatedAt, &bike.UpdatedAt)

		if err != nil {
			return nil, err
		}

		bikes = append(bikes, bike)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bikes, nil
}
