package server

import (
	"io"
	"os"
	"reflect"
	"testing"
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

		store := NewFileSystemStore(database)

		got, err := store.GetLeague()

		if err != nil {
			t.Fatal(err)
		}
		want := League{
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

		store := NewFileSystemStore(database)
		var err error
		want := League{
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

func TestFileSystemStore_GetScore(t *testing.T) {
	t.Parallel()

	database, cleanDatabase := createTempFile(t, initialData)
	defer cleanDatabase()

	store := NewFileSystemStore(database)

	got, err := store.GetPlayerScore("Pepper")

	if err != nil {
		t.Fatal(err)
	}

	want := 2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFileSystemStore_ScoreWin(t *testing.T) {
	t.Run("score win for existing player", func(t *testing.T) {
		t.Parallel()

		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()
		store := NewFileSystemStore(database)

		var err error
		_, err = store.GetPlayerScore("Pepper")
		if err != nil {
			t.Fatal(err)
		}

		store.ScoreWin("Pepper")

		got, err := store.GetPlayerScore("Pepper")
		if err != nil {
			t.Fatal(err)
		}
		want := 3
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("score win for new player", func(t *testing.T) {
		t.Parallel()

		database, cleanDatabase := createTempFile(t, initialData)
		defer cleanDatabase()
		store := NewFileSystemStore(database)

		var err error

		store.ScoreWin("Sapo")
		store.ScoreWin("Sapo")

		got, err := store.GetPlayerScore("Sapo")
		if err != nil {
			t.Fatal(err)
		}
		want := 2
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
