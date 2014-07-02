package main

import (
	"github.com/jtwb/jason-learns-go/server2/app/handlers"
	"net/http"
)

func main() {
	h := handlers.NewStatus()
	http.Handle("/status", h)
	http.ListenAndServe("localhost:4000", nil)
}
