package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/isaporiti/learn-go-with-tests/server/server"
	store "github.com/isaporiti/learn-go-with-tests/server/store"
)

func TestGetPlayers(t *testing.T) {
	playerStore := store.NewInMemoryPlayerStore(map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	})
	playerServer := server.NewPlayerServer(playerStore)
	t.Run("it returns Pepper score", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, http.MethodGet, "/players/Pepper")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertEqual(t, "20", response.Body.String())
	})

	t.Run("it returns Floyd's score", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, http.MethodGet, "/players/Floyd")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertEqual(t, "10", response.Body.String())
	})

	t.Run("it informs if player is not found", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, http.MethodGet, "/players/Unkown")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertEqual(t, http.StatusNotFound, response.Code)
		assertEqual(t, "", response.Body.String())
	})
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
		playerStore := store.NewInMemoryPlayerStore(map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		})
		playerServer := server.NewPlayerServer(playerStore)
		request := newRequest(t, http.MethodPost, "/players/Pepper")
		response := httptest.NewRecorder()

		playerServer.ServeHTTP(response, request)

		assertEqual(t, http.StatusAccepted, response.Code)
	})
}
