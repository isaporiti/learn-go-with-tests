package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/isaporiti/learn-go-with-tests/server"
)

func TestGetPlayers(t *testing.T) {
	t.Run("it returns Pepper score", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, "/players/Pepper")
		response := httptest.NewRecorder()

		server.PlayerServer(response, request)

		assertEqual(t, "20", response.Body.String())
	})

	t.Run("it returns Floyd's score", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, "/players/Floyd")
		response := httptest.NewRecorder()

		server.PlayerServer(response, request)

		assertEqual(t, "10", response.Body.String())
	})

	t.Run("it informs if player is not found", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, "/players/Unkown")
		response := httptest.NewRecorder()

		server.PlayerServer(response, request)

		assertEqual(t, http.StatusNotFound, response.Code)
		assertEqual(t, "", response.Body.String())
	})
}

func newRequest(t *testing.T, path string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, path, nil)
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
