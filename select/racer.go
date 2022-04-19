package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	timeA := time.Now()
	http.Get(a)
	durationA := time.Since(timeA)

	timeB := time.Now()
	http.Get(a)
	durationB := time.Since(timeB)

	if durationA < durationB {
		return a
	}
	return b
}
