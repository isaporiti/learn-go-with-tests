package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	playerStore PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) PlayerServer {
	s := PlayerServer{}
	s.playerStore = store
	router := http.NewServeMux()
	router.HandleFunc("/league", s.handleLeague)
	router.HandleFunc("/players/", s.handlePlayers)
	s.Handler = router
	return s
}

func (s *PlayerServer) handleLeague(response http.ResponseWriter, request *http.Request) {
	leagueTable, _ := s.getLeagueTable()
	response.Header().Set("content-type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(leagueTable)
}

func (s *PlayerServer) getLeagueTable() ([]Player, error) {
	return s.playerStore.GetLeague()
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

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	GetLeague() ([]Player, error)
	ScoreWin(name string) error
}

type Player struct {
	Name string
	Wins int
}
