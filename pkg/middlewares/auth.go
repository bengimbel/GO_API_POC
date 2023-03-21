package middlewares

import (
	"net/http"
	"strings"

	"github.com/bengimbel/go-bookstore/pkg/auth"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		bearerToken := strings.Split(tokenString, " ")
		err := auth.ValidateToken(bearerToken[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		// name := claims.(jwt.MapClaims)["name"].(string)
		// role := claims.(jwt.MapClaims)["role"].(string)

		// r.Header.Set("name", name)
		// r.Header.Set("role", role)

		next.ServeHTTP(w, r)
	})
}
