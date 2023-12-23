package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/mail"
	"time"

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

	hashedPassword, err := s.authStore.GetHashedPassword(creds.Email)
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

	refreshToken, err := services.NewRefreshToken(creds.Email)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Expires:  time.Now().Add(72 * time.Hour), 
		HttpOnly: true,                           
		Path:     "/",                            
		Secure:   true,                           
		SameSite: http.SameSiteStrictMode,        
	})

	return helpers.WriteJSON(w, http.StatusOK, LoginDTO{token})
}

func (s *APIServer) handleRefresh(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie("refreshToken")
	if err != nil {
		if err == http.ErrNoCookie {
			return fmt.Errorf("no refresh token cookie found")
		}
		return err
	}

	email, err := services.VerifyJWTToken(cookie.Value)
	if err != nil {
		return err
	}

	newAccessToken, err := services.NewAccessToken(email)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, LoginDTO{newAccessToken})
}
