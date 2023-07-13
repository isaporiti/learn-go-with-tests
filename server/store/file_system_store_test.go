package store_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/isaporiti/learn-go-with-tests/server/server"
	"github.com/isaporiti/learn-go-with-tests/server/store"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get league", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Pepper", "Wins": 2},
			{"Name": "Floyd", "Wins": 3}
		]`)
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
		database := strings.NewReader(`[
			{"Name": "Pepper", "Wins": 2},
			{"Name": "Floyd", "Wins": 3}
		]`)
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
