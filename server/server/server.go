package server

import (
	"fmt"
	"net/http"
	"strings"

	store "github.com/isaporiti/learn-go-with-tests/server/store"
)

type PlayerServer struct {
	playerStore store.PlayerStore
}

func NewPlayerServer(store store.PlayerStore) PlayerServer {
	return PlayerServer{store}
}

func (server *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		response.WriteHeader(http.StatusAccepted)
		return
	}
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	score, err := server.playerStore.GetPlayerScore(player)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprint(response, "")
		return
	}
	fmt.Fprint(response, score)
}
