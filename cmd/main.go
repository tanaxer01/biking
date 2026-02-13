package main

import (
	"github.com/tanaxer01/biking/internal/infra/http"
	"github.com/tanaxer01/biking/internal/infra/sqlite"
)

func main() {
	// TODO: Config things the right way
	_, err := sqlite.NewBikingDB("biking.db")
	if err != nil {
		panic(err)
	}

	server := http.NewServer(":8080")
	defer server.Close()

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
