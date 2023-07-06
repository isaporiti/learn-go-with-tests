package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(response http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")
	score := getScore(player)
	if score == "" {
		response.WriteHeader(http.StatusNotFound)
		fmt.Fprint(response, "")
		return
	}
	fmt.Fprint(response, score)
}

func getScore(player string) string {
	if player == "Pepper" {
		return "20"
	}
	if player == "Floyd" {
		return "10"
	}
	return ""
}
