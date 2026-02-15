package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/tanaxer01/biking/internal/core/user"
	"github.com/tanaxer01/biking/pkg/biking"
)

type UserHandler struct {
	service *user.Service
}

func NewUserHandler(service *user.Service) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input biking.InsertUser

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Insert(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("OK"))
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement login user
}

func (h *UserHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement get profile
}

func (h *UserHandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement update profile
}
