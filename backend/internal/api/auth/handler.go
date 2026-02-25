package auth

import (
	jwtutil "backend/internal/pkg/jwt-util"
	"backend/internal/pkg/password"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	authRepo *AuthRepo
}

func NewAuthHandler(authRepo *AuthRepo) *AuthHandler {
	return &AuthHandler{
		authRepo: authRepo,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	defer r.Body.Close()

	var body LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid json body", http.StatusInternalServerError)
		return
	}

	user, err := h.authRepo.GetUserByEmail(body.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Invalid Email", http.StatusBadRequest)
		} else {
			http.Error(w, fmt.Sprintf("SQL: %s", err.Error()), http.StatusInternalServerError)
		}
		return
	}

	passwordMatched, err := password.VerifyPassword(body.Password, user.PasswordHash)

	if err != nil {
		http.Error(w, fmt.Sprintf("PASSWORD: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	if !passwordMatched {
		http.Error(w, "Incorrect Credential", http.StatusBadRequest)
		return
	}

	refreshToken, _, err := jwtutil.CreateRefreshToken(user.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ACCESS TOKEN
	token, err := jwtutil.CreateToken(jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iss":     "what-the-fack",
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("JWT: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Path:     "/auth/refresh",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   60 * 60 * 24 * 14,
	})

	response := map[string]string{
		"token": *token,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (h *AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_, err := r.Cookie("refresh_token")

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, ok := r.Context().Value("claims").(*jwtutil.AccessTokenClaims)

	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID := claims.UserID

	accessToken, err := jwtutil.CreateToken(jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iss":     "what-the-fack",
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"access_token": accessToken,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *AuthHandler) TestFunction(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"message": "hello from auth module",
	}

	json.NewEncoder(w).Encode(response)
}
