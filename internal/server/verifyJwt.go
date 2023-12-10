package server

import (
	"net/http"
	"strings"

	"github.com/petar-dobre/gobank/internal/services"
)

func (s *APIServer) VerifyJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		onlyToken := strings.Split(tokenString, " ")[1]

		_, err := services.VerifyJWTToken(onlyToken)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
