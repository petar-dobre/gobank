package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"

	"github.com/petar-dobre/gobank/internal/helpers"
	"github.com/petar-dobre/gobank/internal/services"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginDTO struct {
	AccessToken string `json:"accessToken"`
}

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		return err
	}

	if _, err := mail.ParseAddress(creds.Email); err != nil {
		return err
	}

	hashedPassword, err := s.store.GetHashedPassword(creds.Email)
	if err != nil {
		return err
	}

	if match := helpers.CheckPassword(hashedPassword, creds.Password); !match {
		return fmt.Errorf("passwords do not match")
	}

	token, err := services.NewAccessToken(creds.Email)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, LoginDTO{token})
}
