package store_test

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	server "github.com/isaporiti/learn-go-with-tests/server/server"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("get league", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Pepper", "Wins": 2},
			{"Name": "Floyd", "Wins": 3}
		]`)
		store := FileSystemStore{database}

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
		store := FileSystemStore{database}
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

type FileSystemStore struct {
	database io.ReadSeeker
}

func (s *FileSystemStore) GetLeague() ([]server.Player, error) {
	s.database.Seek(0, 0)
	var league []server.Player
	err := json.NewDecoder(s.database).Decode(&league)
	if err != nil {
		return nil, fmt.Errorf("couln't decode league: %v", err)
	}
	return league, nil
}
