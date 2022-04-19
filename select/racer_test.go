package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := buildDelayedServer(20 * time.Millisecond)
	fastServer := buildDelayedServer(0 * time.Millisecond)
	defer slowServer.Close()
	defer fastServer.Close()

	got := Racer(slowServer.URL, fastServer.URL)
	want := fastServer.URL

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}

func buildDelayedServer(delay time.Duration) *httptest.Server {
	slowServer := httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
	return slowServer
}
