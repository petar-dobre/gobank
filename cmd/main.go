package main

import (
	"log"

	"github.com/lpernett/godotenv"
	"github.com/petar-dobre/gobank/internal/server"
	"github.com/petar-dobre/gobank/internal/store"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	pgStore, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := pgStore.Init(); err != nil {
		log.Fatal(err)
	}

	accountStore := &store.AccountStore{PostgresStore: pgStore}
	authStore := &store.AuthStore{PostgresStore: pgStore}

	server := server.NewAPIServer(":8080", accountStore, authStore)

	server.Run()
}
