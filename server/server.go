package server

import (
	"fmt"
	"net/http"
)

func PlayerServer(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "20")
}
