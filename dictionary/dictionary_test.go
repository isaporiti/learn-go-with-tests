package dictionary

import "testing"

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
	dictionary := Dictionary{}

	dictionary.Add("test", "this is just a test")

	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
