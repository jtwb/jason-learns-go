package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jtwb/jason-learns-go/server2/app/models"
	"net/http"
)

type Status struct {
	Model *models.Status
}

func NewStatus() *Status {
	return &Status{
		Model: models.NewStatus(),
	}
}

func (s Status) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(s.Model)
	if err != nil {
		http.Error(w, "I die", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%s", json)
}
