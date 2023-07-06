package main

import (
	"log"
	"net/http"

	server "github.com/isaporiti/learn-go-with-tests/server"
)

func main() {
	handler := http.HandlerFunc(server.PlayerServer)
	err := http.ListenAndServe(":5001", handler)
	log.Fatal(err)
}
