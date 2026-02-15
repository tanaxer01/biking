package main

import (
	"github.com/tanaxer01/biking/internal/core/admin"
	"github.com/tanaxer01/biking/internal/core/user"
	"github.com/tanaxer01/biking/internal/infra/auth"
	"github.com/tanaxer01/biking/internal/infra/crypto"
	"github.com/tanaxer01/biking/internal/infra/http"
	"github.com/tanaxer01/biking/internal/infra/sqlite"
)

func main() {
	// TODO: Config things the right way
	db, err := sqlite.NewBikingDB("biking.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepo := sqlite.NewUserRepository(db)
	bikeRepo := sqlite.NewBikeRepository(db)
	authSvc := auth.NewAuth("dev-secret")

	userService := user.NewService(userRepo, &crypto.Crypto{}, authSvc)
	userHandler := http.NewUserHandler(userService)

	adminService := admin.NewService(bikeRepo, userRepo)
	adminHandler := http.NewAdminHandler(adminService)

	server := http.NewServer(":8080", userHandler, adminHandler)
	defer server.Close()

	err = server.Start()
	if err != nil {
		panic(err)
	}
}
