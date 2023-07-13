package store_test

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/isaporiti/learn-go-with-tests/server/server"
	"github.com/isaporiti/learn-go-with-tests/server/store"
)

const initialData = `[
	{"Name": "Pepper", "Wins": 2},
	{"Name": "Floyd", "Wins": 3}
]`

func TestFileSystemStore_GetLeague(t *testing.T) {
	t.Run("get league", func(t *testing.T) {
		t.Parallel()
		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

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

		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

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

		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()

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
