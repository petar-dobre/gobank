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

	store, err := store.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := server.NewAPIServer(":8080", store)

	server.Run()
}
