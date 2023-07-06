package main

import (
	"log"
	"net/http"

	server "github.com/isaporiti/learn-go-with-tests/server/server"
	store "github.com/isaporiti/learn-go-with-tests/server/store"
)

func main() {
	store := store.NewInMemoryPlayerStore(map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	})
	server := server.NewPlayerServer(store)
	err := http.ListenAndServe(":5001", &server)
	log.Fatal(err)
}
