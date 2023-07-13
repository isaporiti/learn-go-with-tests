package store_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/isaporiti/learn-go-with-tests/server/server"
	"github.com/isaporiti/learn-go-with-tests/server/store"
)

var database = strings.NewReader(`[
	{"Name": "Pepper", "Wins": 2},
	{"Name": "Floyd", "Wins": 3}
]`)

func TestFileSystemStore_GetLeague(t *testing.T) {
	t.Run("get league", func(t *testing.T) {
		t.Parallel()

		store := store.NewFileSystemStore(database)

		got, err := store.GetLeague()

		if err != nil {
			t.Fatal(err)
		}
		want := []server.Player{
			{Name: "Pepper", Wins: 2},
			{Name: "Floyd", Wins: 3},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got league: %v, want: %v", got, want)
		}
	})

	t.Run("get league multiple times", func(t *testing.T) {
		t.Parallel()

		store := store.NewFileSystemStore(database)
		var err error
		want := []server.Player{
			{Name: "Pepper", Wins: 2},
			{Name: "Floyd", Wins: 3},
		}

		_, err = store.GetLeague()
		if err != nil {
			t.Fatal(err)
		}
		got, err := store.GetLeague()
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got league: %v, want: %v", got, want)
		}
	})
}

func TestFileSystemStore_GetScore(t *testing.T) {
	t.Run("get score", func(t *testing.T) {
		t.Parallel()

		store := store.NewFileSystemStore(database)

		got, err := store.GetPlayerScore("Pepper")

		if err != nil {
			t.Errorf("could not get score: %v", err)
		}

		want := 2
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
