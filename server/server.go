package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(response http.ResponseWriter, request *http.Request) {
	player := strings.TrimPrefix(request.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprint(response, "20")
		return
	}
	if player == "Floyd" {
		fmt.Fprint(response, "10")
		return
	}
	response.WriteHeader(http.StatusNotFound)
	fmt.Fprint(response, "")
}
