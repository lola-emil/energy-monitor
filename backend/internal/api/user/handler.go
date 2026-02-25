package user

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	userRepo *UserRepo
}

func NewUserHandler(userRepo *UserRepo) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	response := map[string]any{
		"message": "Hello World",
	}

	json.NewEncoder(w).Encode(response)
}
