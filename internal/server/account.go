package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/petar-dobre/gobank/internal/helpers"
	"github.com/petar-dobre/gobank/internal/models"
)

func (s *APIServer) handleGetAccounts(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.accountStore.GetAccounts()
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, accounts)
}

func (s *APIServer) handleGetAccountByID(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetID(r)
	if err != nil {
		return err
	}

	account, err := s.accountStore.GetAccountByID(id)
	if err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	type CreateAccountRequest struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	createAccountReq := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(createAccountReq); err != nil {
		return err
	}

	hashedPassword, err := helpers.HashPassword(createAccountReq.Password)
	if err != nil {
		return err
	}

	account := models.NewAccount(createAccountReq.FirstName,
		createAccountReq.LastName,
		createAccountReq.Email,
		hashedPassword,
	)

	if err := s.accountStore.CreateAccount(account); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) hanldeUpdateAccount(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		return fmt.Errorf("invalid id given %s", idStr)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("invalid id given %s", idStr)
	}

	updateAccountReq := new(models.UpdateAccountReq)

	if err := json.NewDecoder(r.Body).Decode(updateAccountReq); err != nil {
		return err
	}

	if err := s.accountStore.UpdateAccount(updateAccountReq, id); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, map[string]int{"updated": id})
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	id, err := helpers.GetID(r)
	if err != nil {
		return err
	}

	if err := s.accountStore.DeleteAccount(id); err != nil {
		return err
	}

	return helpers.WriteJSON(w, http.StatusOK, map[string]int{"deleted": id})
}
