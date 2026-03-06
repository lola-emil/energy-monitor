package user

import (
	"backend/internal/pkg/password"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
	pathUserId := r.PathValue("id")
	id, err := strconv.ParseInt(pathUserId, 10, 64)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetUserById(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Response error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.GetUsers(r.Context())

	if err != nil {
		log.Println("DB error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println("Response error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	var body UpdateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid json body", http.StatusInternalServerError)
		return
	}

	pathUserId := r.PathValue("id")
	id, err := strconv.ParseInt(pathUserId, 10, 64)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if body.Password != nil {
		hashedPassword, err := password.HashPassword(*body.Password, password.DefaultParams)

		if err != nil {
			http.Error(w, "Password error", http.StatusInternalServerError)
			return
		}

		body.Password = &hashedPassword
	}

	rowsAffected, err := h.userRepo.UpdateUserById(r.Context(), id, body)

	if err != nil {
		log.Println("DB error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"message":       "Updated successful",
		"rows_affected": rowsAffected,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Response error:", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
