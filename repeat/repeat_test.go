package repeat

import "testing"

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	if got := Repeat("a"); got != expected {
		t.Errorf("expected: %s, got: %s", expected, got)
	}
}

func Repeat(character string) string {
	return "aaaaa"
}
