package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "Ignacio")

	got := buffer.String()
	want := "Hello, Ignacio"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
