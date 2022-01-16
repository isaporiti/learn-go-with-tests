package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, err := dictionary.Search("test")
		want := "this is just a test"

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
		if err != nil {
			t.Errorf("got error %q want nil given, %q", got, "test")
		}
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("unknown")
		want := wordNotFoundError.Error()
		got := err.Error()

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "unknown")
		}
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"
		dictionary.Add(word, definition)

		want := definition
		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("should find added word:", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		got := dictionary.Add(word, definition)

		want := wordAlreadyDefinedError

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "this is a test"}
		newDefinition := "this is the new definition"

		dictionary.Update(word, newDefinition)

		want := newDefinition
		got, _ := dictionary.Search(word)

		if got != want {
			t.Errorf("got %q want %q given %q", got, want, word)
		}
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}
		definition := "this is a test"

		got := dictionary.Update(word, definition)

		want := wordDoesNotExistError

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	dictionary.Add(word, "this is a test")

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != wordNotFoundError {
		t.Errorf("expected %q to be deleted", word)
	}
}
