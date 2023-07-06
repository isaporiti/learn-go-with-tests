package server

import (
	"fmt"
	"net/http"
	"strings"

	store "github.com/isaporiti/learn-go-with-tests/server/store"
)

type PlayerServer struct {
	playerStore store.PlayerStore
	router      *http.ServeMux
}

func NewPlayerServer(store store.PlayerStore) PlayerServer {
	s := PlayerServer{store, http.NewServeMux()}
	s.router.HandleFunc("/league", s.handleLeague)
	s.router.HandleFunc("/players/", s.handlePlayers)
	return s
}

func (s *PlayerServer) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(response, request)
}

func (s *PlayerServer) handleLeague(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}

func (s *PlayerServer) handlePlayers(response http.ResponseWriter, request *http.Request) {
	player := s.getPlayerName(request)
	if request.Method == http.MethodPost {
		s.scoreWin(player)
		response.WriteHeader(http.StatusAccepted)
		return
	}
	if request.Method == http.MethodGet {
		score, err := s.getScore(player)
		if err != nil {
			response.WriteHeader(http.StatusNotFound)
			fmt.Fprint(response, "")
			return
		}
		fmt.Fprint(response, score)
	}
}
func (*PlayerServer) getPlayerName(request *http.Request) string {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	return player
}

func (s *PlayerServer) scoreWin(player string) {
	s.playerStore.ScoreWin(player)
}

func (s *PlayerServer) getScore(player string) (int, error) {
	return s.playerStore.GetPlayerScore(player)
}
