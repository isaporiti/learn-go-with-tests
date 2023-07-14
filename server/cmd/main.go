package main

import (
	"log"
	"net/http"
	"os"

	"github.com/isaporiti/learn-go-with-tests/server"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	db.Write([]byte("[]"))
	playerStore := server.NewFileSystemStore(db)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewPlayerServer(playerStore)
	log.Fatal(http.ListenAndServe(":5001", &server))
}
