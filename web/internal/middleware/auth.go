package appmiddleware

import (
	"context"
	jwtutil "energy-monitor-server/internal/utils/jwt"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		log.Println("HIT")

		if authHeader == "" {
			log.Println("Missing Token")
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		log.Println("HIT1")

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwtutil.Parse(tokenStr)
		if err != nil {
			log.Println("INVALID TOKEN")
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		log.Println("HIT2")

		ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
		ctx = context.WithValue(ctx, "role", claims.Role)

		log.Println("HIT3")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
