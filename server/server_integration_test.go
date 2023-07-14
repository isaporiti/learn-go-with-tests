package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	t.Parallel()
	database, clearDatabase := createTempFile(t, `[
		{"Name": "Pepper", "Wins": 20},
		{"Name": "Floyd", "Wins": 10}
	]`)
	defer clearDatabase()
	store := NewFileSystemStore(database)
	playerServer := NewPlayerServer(store)
	scoreWin(t, playerServer, "Pepper")
	scoreWin(t, playerServer, "Pepper")
	response := getScore(t, playerServer, "Pepper")

	assertEqual(t, "22", response.Body.String())
}

func TestRecordingWinsAndRetrievingLeague(t *testing.T) {
	t.Parallel()
	var err error
	database, clearDatabase := createTempFile(t, `[
		{"Name": "Pepper", "Wins": 20},
		{"Name": "Floyd", "Wins": 10}
	]`)
	defer clearDatabase()
	store := NewFileSystemStore(database)
	playerServer := NewPlayerServer(store)
	scoreWin(t, playerServer, "Pepper")
	scoreWin(t, playerServer, "Floyd")
	scoreWin(t, playerServer, "Floyd")
	scoreWin(t, playerServer, "Floyd")

	request := newRequest(t, http.MethodGet, "/league")
	response := httptest.NewRecorder()

	playerServer.ServeHTTP(response, request)

	var got League
	want := League{
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
