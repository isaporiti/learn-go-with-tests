package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("it races HTTP requests between two URLs and tells the winner", func(t *testing.T) {
		slowServer := buildDelayedServer(5 * time.Millisecond)
		fastServer := buildDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		got, _ := Racer(slowServer.URL, fastServer.URL, 10*time.Millisecond)
		want := fastServer.URL

		if want != got {
			t.Errorf("want %s, got %s", want, got)
		}
	})

	t.Run("it informs when both requests took more than 10 seconds", func(t *testing.T) {
		slowServer := buildDelayedServer(12 * time.Millisecond)
		fastServer := buildDelayedServer(11 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("wanted an error, got nil")
		}
	})
}

func buildDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}
