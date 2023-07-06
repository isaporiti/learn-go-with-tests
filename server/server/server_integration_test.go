package server_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	server "github.com/isaporiti/learn-go-with-tests/server/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
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
}

func TestRecordingWinsAndRetrievingLeague(t *testing.T) {
	t.Parallel()
	var err error
	playerStore, err := newPlayerStore()
	if err != nil {
		t.Error(err)
	}
	playerServer := server.NewPlayerServer(playerStore)
	scoreWin(t, playerServer, "Pepper")
	scoreWin(t, playerServer, "Floyd")
	scoreWin(t, playerServer, "Floyd")
	scoreWin(t, playerServer, "Floyd")

	request := newRequest(t, http.MethodGet, "/league")
	response := httptest.NewRecorder()

	playerServer.ServeHTTP(response, request)

	var got []server.Player
	want := []server.Player{
		{Name: "Pepper", Wins: 21},
		{Name: "Floyd", Wins: 13},
	}
	err = json.NewDecoder(response.Body).Decode(&got)
	if err != nil {
		t.Errorf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
	}
	assertEqual(t, http.StatusOK, response.Code)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: '%v', want: '%v'", got, want)
	}
}
