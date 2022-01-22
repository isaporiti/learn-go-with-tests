package countdown

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{0}
	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to Sleeper: want 4, got %d", spySleeper.Calls)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}
