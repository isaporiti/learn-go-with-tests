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

func (s *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodPost {
		s.scoreWin(request, response)
		return
	}
	if request.Method == http.MethodGet {
		s.getScore(request, response)
	}
}

func (s *PlayerServer) scoreWin(request *http.Request, response http.ResponseWriter) {
	player := s.getPlayerName(request)
	s.playerStore.ScoreWin(player)
	response.WriteHeader(http.StatusAccepted)
}

func (*PlayerServer) getPlayerName(request *http.Request) string {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	return player
}

func (s *PlayerServer) getScore(request *http.Request, response http.ResponseWriter) {
	player := s.getPlayerName(request)
	score, err := s.playerStore.GetPlayerScore(player)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprint(response, "")
		return
	}
	fmt.Fprint(response, score)
}
