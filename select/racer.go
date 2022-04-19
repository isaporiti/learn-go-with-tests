package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	durationA := measureTime(a)
	durationB := measureTime(b)
	if durationA < durationB {
		return a
	}
	return b
}

func measureTime(a string) time.Duration {
	timeA := time.Now()
	http.Get(a)
	durationA := time.Since(timeA)
	return durationA
}
