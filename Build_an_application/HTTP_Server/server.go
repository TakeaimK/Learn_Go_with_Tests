package main

import (
	"fmt"
	"net/http"
	"strings"
)

type Handler interface {
	ServerHTTP(http.ResponseWriter, *http.Request)
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {

	// TrimPrefix returns s without the provided leading prefix string.
	// If s doesn't start with prefix, s is returned unchanged.
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	if player == "Pepper" {
		fmt.Fprint(w, "20")
		return
	}

	if player == "Floyd" {
		fmt.Fprint(w, "10")
		return
	}
}
