package validationutil

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func WriteValidationErrors(w http.ResponseWriter, err error) {
	validationErrors := err.(validator.ValidationErrors)

	errors := make(map[string]string)

	for _, e := range validationErrors {
		switch e.Tag() {
		case "required":
			errors[e.Field()] = "is required"
		case "email":
			errors[e.Field()] = "must be a valid email"
		case "min":
			errors[e.Field()] = "must be at least " + e.Param() + " characters"
		default:
			errors[e.Field()] = "is invalid"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]any{
		"errors": errors,
	})
}
