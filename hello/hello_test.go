package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, World"

	got := Hello("", "")

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestHelloWithName(t *testing.T) {
	expected := "Hello, Ignacio"

	got := Hello("Ignacio", "")

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func TestHelloInSpanish(t *testing.T) {
	expected := "Hola, Ignacio"

	got := Hello("Ignacio", "Spanish")

	if got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}
