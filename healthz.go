package main

import (
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	reponsdWithJSON(w, 200, struct{}{})
}
