package server_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/isaporiti/learn-go-with-tests/server/server"
	store "github.com/isaporiti/learn-go-with-tests/server/store"
)

func TestGetPlayers(t *testing.T) {
	playerStore, err := newPlayerStore()
	if err != nil {
		t.Error(err)
	}
	playerServer := server.NewPlayerServer(playerStore)

	t.Run("it returns Pepper score", func(t *testing.T) {
		t.Parallel()
		response := getScore(t, playerServer, "Pepper")

		assertEqual(t, "20", response.Body.String())
	})

	t.Run("it returns Floyd's score", func(t *testing.T) {
		t.Parallel()
		response := getScore(t, playerServer, "Floyd")

		assertEqual(t, "10", response.Body.String())
	})

	t.Run("it informs if player is not found", func(t *testing.T) {
		t.Parallel()
		response := getScore(t, playerServer, "Unknown")

		assertEqual(t, http.StatusNotFound, response.Code)
		assertEqual(t, "", response.Body.String())
	})
}

func newPlayerStore() (*store.InMemoryPlayerStore, error) {
	scores := map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}
	playerStore, err := store.NewInMemoryPlayerStore(store.WithScores(scores))
	return playerStore, err
}

func getScore(t *testing.T, playerServer server.PlayerServer, player string) *httptest.ResponseRecorder {
	path := fmt.Sprintf("/players/%s", player)
	request := newRequest(t, http.MethodGet, path)
	response := httptest.NewRecorder()

	playerServer.ServeHTTP(response, request)
	return response
}

func newRequest(t *testing.T, method, path string) *http.Request {
	request, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Errorf("couldn't create request: %s", err)
	}
	return request
}

func assertEqual[T comparable](t *testing.T, want, got T) {
	if want != got {
		t.Errorf("want '%v', got '%v'", want, got)
	}
}

func TestPostPlayer(t *testing.T) {
	t.Run("it returns Accepted code", func(t *testing.T) {
		t.Parallel()
		playerStore, err := newPlayerStore()
		if err != nil {
			t.Error(err)
		}
		playerServer := server.NewPlayerServer(playerStore)
		response := scoreWin(t, playerServer, "Pepper")

		assertEqual(t, http.StatusAccepted, response.Code)
	})

	t.Run("it records win", func(t *testing.T) {
		t.Parallel()
		playerStore, err := newPlayerStore()
		if err != nil {
			t.Error(err)
		}
		playerServer := server.NewPlayerServer(playerStore)
		scoreWin(t, playerServer, "Pepper")
		scoreWin(t, playerServer, "Pepper")
		response := getScore(t, playerServer, "Pepper")

		assertEqual(t, "22", response.Body.String())
	})
}

func scoreWin(t *testing.T, playerServer server.PlayerServer, player string) *httptest.ResponseRecorder {
	path := fmt.Sprintf("/players/%s", player)
	request := newRequest(t, http.MethodPost, path)
	response := httptest.NewRecorder()
	playerServer.ServeHTTP(response, request)
	return response
}

func TestLeague(t *testing.T) {
	playerStore, err := newPlayerStore()
	if err != nil {
		t.Error(err)
	}
	playerServer := server.NewPlayerServer(playerStore)

	t.Run("it returns OK on /league", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, http.MethodGet, "/league")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		var got []server.Player
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertEqual(t, http.StatusOK, response.Code)
	})
}

