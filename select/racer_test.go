package racer

import "testing"

func TestRacer(t *testing.T) {
	slowUrl := "https://www.facebook.com"
	fastUrl := "https://www.quii.dev"

	got := Racer(slowUrl, fastUrl)
	want := fastUrl

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
