package admin

type Service struct {
	db DB
}

type DB interface {
}

func NewService(db DB) *Service {
	return &Service{db: db}
}
