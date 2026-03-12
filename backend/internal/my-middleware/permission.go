package mymiddleware

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleUser  Role = "USER"
)

var permissions = map[Role]map[string]map[string]bool{
	RoleAdmin: {
		"/users": {
			http.MethodGet:    true,
			http.MethodPost:   true,
			http.MethodPut:    true,
			http.MethodDelete: true,
		},
		"/devices": {
			http.MethodGet:    true,
			http.MethodPost:   true,
			http.MethodPut:    true,
			http.MethodDelete: true,
		},
		"/device-claims": {
			http.MethodGet:    true,
			http.MethodPost:   true,
			http.MethodPut:    true,
			http.MethodDelete: true,
		},
	},

	RoleUser: {
		"/users": {
			http.MethodGet: true,
			http.MethodPut: true,
		},
		"/devices": {
			http.MethodGet: true,
		},
		"/device-claims": {
			http.MethodGet:    true,
			http.MethodPost:   true,
			http.MethodPut:    true,
			http.MethodDelete: true,
		},
	},
}

func Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Example: role stored in context by auth middleware
		roleVal := r.Context().Value("role")
		fmt.Println("Role:", roleVal)

		if roleVal == nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		role := Role(roleVal.(string))

		// Get chi route pattern
		routeContext := chi.RouteContext(r.Context())
		path := routeContext.RoutePattern()

		method := r.Method

		rolePerms, ok := permissions[role]
		if !ok {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		pathPerms, ok := rolePerms[path]
		if !ok {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		if !pathPerms[method] {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
