package main

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	if player == "Maria" {
		fmt.Fprint(w, "20")
		return
	}

	if player == "Pedro" {
		fmt.Fprint(w, "10")
		return
	}
}
