package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(response http.ResponseWriter, request *http.Request) {
	if strings.Contains(request.URL.Path, "Pepper") {
		fmt.Fprint(response, "20")
		return
	}
	fmt.Fprint(response, "10")
}
