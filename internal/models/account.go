package models

import (
	"math/rand"
	"net/http"
	"time"
)

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type AccountDTO struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type AccountHandler interface {
	handleGetAccounts(w http.ResponseWriter, r *http.Request) error
	handleGetAccountByID(w http.ResponseWriter, r *http.Request) error
	handleCreateAccount(w http.ResponseWriter, r *http.Request) error
	hanldeUpdateAccount(w http.ResponseWriter, r *http.Request) error
	handleDeleteAccount(w http.ResponseWriter, r *http.Request) error
}

type UpdateAccountReq struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

func NewAccount(firstName, lastName, email, password string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Number:    int64(rand.Intn(100000)),
		CreatedAt: time.Now().UTC(),
	}
}
