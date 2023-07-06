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

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("want '%s', got '%s'", got, want)
		}
	})

	t.Run("it returns Floyd's score", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, "/players/Floyd")
		response := httptest.NewRecorder()

		server.PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("want '%s', got '%s'", got, want)
		}
	})

	t.Run("it informs if player is not found", func(t *testing.T) {
		t.Parallel()
		request := newRequest(t, "/players/Unkown")
		response := httptest.NewRecorder()

		server.PlayerServer(response, request)

		if response.Code != http.StatusNotFound {
			t.Errorf("want '%d', got '%d'", response.Code, http.StatusNotFound)
		}
		if response.Body.String() != "" {
			t.Errorf("want '%s', got '%s'", response.Body.String(), "")
		}
	})
}

func newRequest(t *testing.T, path string) *http.Request {
	request, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		t.Errorf("couldn't create request: %s", err)
	}
	return request
}
