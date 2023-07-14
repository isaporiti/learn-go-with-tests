package server_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/isaporiti/learn-go-with-tests/server/server"
	"github.com/isaporiti/learn-go-with-tests/server/store"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	t.Parallel()
	database, clearDatabase := createTempFile(t, `[
		{"Name": "Pepper", "Wins": 20},
		{"Name": "Floyd", "Wins": 10}
	]`)
	defer clearDatabase()
	store := store.NewFileSystemStore(database)
	playerServer := server.NewPlayerServer(store)
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
	store := store.NewFileSystemStore(database)
	playerServer := server.NewPlayerServer(store)
	scoreWin(t, playerServer, "Pepper")
	scoreWin(t, playerServer, "Floyd")
	scoreWin(t, playerServer, "Floyd")
	scoreWin(t, playerServer, "Floyd")

	request := newRequest(t, http.MethodGet, "/league")
	response := httptest.NewRecorder()

	playerServer.ServeHTTP(response, request)

	var got server.League
	want := server.League{
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

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	var err error
	file, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file: %v", err)
	}
	file.Write([]byte(initialData))
	removeFile := func() {
		err = file.Close()
		if err != nil {
			t.Fatalf("could not close temp file: %v", err)
		}
		os.Remove(file.Name())
	}

	return file, removeFile
}
