package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type BikingDB struct {
	db *sql.DB
}

func NewBikingDB(file string) (*BikingDB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	return &BikingDB{db: db}, nil
}

func (b *BikingDB) Close() error {
	return b.db.Close()
}
