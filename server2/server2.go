package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusHandler struct {
	Status  string
	Message string
	Error   string
}

func (s StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(s)
	if err != nil {
		http.Error(w, "I die", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%s", json)
}

func main() {
	http.Handle("/status", &StatusHandler{"OK", "All good", ""})
	http.ListenAndServe("localhost:4000", nil)
}
