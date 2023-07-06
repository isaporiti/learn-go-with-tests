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
		request, err := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		if err != nil {
			t.Errorf("could't create request: %s", err)
		}
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
		request, err := http.NewRequest(http.MethodGet, "players/Floyd", nil)
		if err != nil {
			t.Errorf("couldn't create request: %s", err)
		}
		response := httptest.NewRecorder()

		server.PlayerServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("want '%s', got '%s'", got, want)
		}
	})
}
