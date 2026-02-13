package biking

import "time"

type RentalData struct {
	UserID         int
	BikeID         int
	StartTime      time.Time
	EndTime        time.Time
	StartLatitude  float64
	StartLongitude float64
	EndLatitude    float64
	EndLongitude   float64
}

type Rental struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	RentalData
}
