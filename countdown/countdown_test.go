package countdown

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	got := buffer.String()
	want := "3"

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
