package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jtwb/jason-learns-go/server2/app/models"
	"net/http"
)

type Status struct {
}

func NewStatus() *Status {
	return new(Status)
}

func (s Status) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	model := models.NewStatus()
	json, err := json.Marshal(model)
	if err != nil {
		http.Error(w, "I die", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%s", json)
}
