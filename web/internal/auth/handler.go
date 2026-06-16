package auth

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type AuthHandler struct {
	service *AuthService
}

func NewAuthHandler(service *AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(AuthResponse{
		Token: token,
	})
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" || req.Name == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	token, err := h.service.Register(r.Context(), req.Username, req.Name, req.Password)
	if err != nil {
		if errors.Is(err, ErrEmailAlreadyExists) {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		log.Println(err.Error())
		http.Error(w, "failed to register", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(RegisterResponse{
		Token: token,
	})
}
