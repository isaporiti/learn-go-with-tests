package main

import (
	"log"
	"net/http"

	server "github.com/isaporiti/learn-go-with-tests/server/server"
	store "github.com/isaporiti/learn-go-with-tests/server/store"
)

func main() {
	playerStore, err := store.NewInMemoryPlayerStore()
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewPlayerServer(playerStore)
	log.Fatal(http.ListenAndServe(":5001", &server))
}
