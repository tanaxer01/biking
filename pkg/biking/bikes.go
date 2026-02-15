package biking

type BikeData struct {
	Available bool    `json:"is_available"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Bike struct {
	ID        int
	CreatedAt string
	UpdatedAt string
	BikeData
}
