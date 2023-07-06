package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	store "github.com/isaporiti/learn-go-with-tests/server/store"
)

type PlayerServer struct {
	playerStore store.PlayerStore
	http.Handler
}

func NewPlayerServer(store store.PlayerStore) PlayerServer {
	s := PlayerServer{}
	s.playerStore = store
	router := http.NewServeMux()
	router.HandleFunc("/league", s.handleLeague)
	router.HandleFunc("/players/", s.handlePlayers)
	s.Handler = router
	return s
}

func (s *PlayerServer) handleLeague(response http.ResponseWriter, request *http.Request) {
	leagueTable := s.getLeagueTable()
	json.NewEncoder(response).Encode(leagueTable)
	response.WriteHeader(http.StatusOK)
}

func (s *PlayerServer) getLeagueTable() []Player {
	var league []Player
	for name, wins := range s.playerStore.GetAllScores() {
		league = append(league, Player{Name: name, Wins: wins})
	}
	return league
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

type Player struct {
	Name string
	Wins int
}
