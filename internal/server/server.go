package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/petar-dobre/gobank/internal/helpers"
	"github.com/petar-dobre/gobank/internal/store"
)

type APIServer struct {
	listenAddr string
	store      *store.PostgresStore
	router     *mux.Router
}

func NewAPIServer(listenAddr string, store *store.PostgresStore) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
		router:     mux.NewRouter(),
	}
}

func (s *APIServer) routes() {
	s.router.HandleFunc("/login", helpers.MakeHTTPHandleFunc(s.handleLogin))
	s.router.HandleFunc("/account", s.VerifyJWT(helpers.MakeHTTPHandleFunc(s.handleAccount)))
	s.router.HandleFunc("/account/{id}", s.VerifyJWT(helpers.MakeHTTPHandleFunc(s.handleGetAccountByID)))
}

func (s *APIServer) Run() {
	s.routes()

	log.Print("Server started running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, s.router)
}
