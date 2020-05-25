package main

import (
	"fmt"
	"net/http"
)

type Handler interface {
	ServerHTTP(http.ResponseWriter, *http.Request)
}

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "20")
}
