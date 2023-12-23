package models

import "net/http"


type AuthHandler interface {
	handleLogin(w http.ResponseWriter, r *http.Request) error
	handleRefresh(w http.ResponseWriter, r *http.Request) error
}