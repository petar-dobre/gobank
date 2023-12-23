package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petar-dobre/gobank/internal/helpers"
	"github.com/petar-dobre/gobank/internal/middleware"
	"github.com/petar-dobre/gobank/internal/models"
)

type APIServer struct {
	listenAddr string
	accountStore      models.AccountStorer
	authStore models.AuthStorer
	router     *mux.Router
}

func NewAPIServer(listenAddr string, accountStore models.AccountStorer, authStore models.AuthStorer) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		accountStore:      accountStore,
		authStore: authStore,
		router:     mux.NewRouter(),
	}
}

func (s *APIServer) routes() {
	// auth
	s.router.HandleFunc("/login", helpers.MakeHTTPHandleFunc(s.handleLogin))
	s.router.HandleFunc("/refresh", helpers.MakeHTTPHandleFunc(s.handleRefresh))

	// account
	s.router.HandleFunc("/account", middleware.VerifyJWT(helpers.MakeHTTPHandleFunc(s.handleGetAccounts))).Methods("GET")
	s.router.HandleFunc("/account", middleware.VerifyJWT(helpers.MakeHTTPHandleFunc(s.handleCreateAccount))).Methods("POST")
	s.router.HandleFunc("/account", middleware.VerifyJWT(helpers.MakeHTTPHandleFunc(s.hanldeUpdateAccount))).Methods("PATCH")
	s.router.HandleFunc("/account", middleware.VerifyJWT(helpers.MakeHTTPHandleFunc(s.handleDeleteAccount))).Methods("DELETE")

	// account/id
	s.router.HandleFunc("/account/{id}", middleware.VerifyJWT(helpers.MakeHTTPHandleFunc(s.handleGetAccountByID)))
}

func (s *APIServer) Run() {
	s.routes()

	log.Print("Server started running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, s.router)
}
