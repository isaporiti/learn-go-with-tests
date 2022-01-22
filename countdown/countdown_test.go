package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("countdown down from 3", func(t *testing.T) {
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
	})

	t.Run("sleep before every print", func(t *testing.T) {
		operations := &SpyCountdownOperations{}
		want := []string{
			sleep, write,
			sleep, write,
			sleep, write,
			sleep, write,
		}

		Countdown(operations, operations)

		if !reflect.DeepEqual(operations.Calls, want) {
			t.Errorf("wanted calls %v got %v", want, operations.Calls)
		}
	})
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

const sleep = "sleep"

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

const write = "write"

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}
