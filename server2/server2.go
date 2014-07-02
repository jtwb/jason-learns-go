package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"syscall"
)

type StatusHandler struct {
	Model *StatusModel
}

type StatusModel struct {
	Status   string
	Message  string
	Hostname string
	Uname    *HappyUtsname
	Errors   []string
}

type HappyUtsname struct {
	parent syscall.Utsname `json:",string"`
}

func NewStatusModel() *StatusModel {
	var (
		errors []string
		uname  syscall.Utsname
	)

	hostname, err := os.Hostname()
	if err != nil {
		errors = append(errors, err.Error())
	}

	err = syscall.Uname(&uname)
	if err != nil {
		errors = append(errors, err.Error())
	}

	return &StatusModel{
		Status:   "OK",
		Hostname: hostname,
		Uname:    &HappyUtsname{uname},
		Errors:   errors,
	}
}

func (s StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Model = NewStatusModel()
	json, err := json.Marshal(s.Model)
	if err != nil {
		http.Error(w, "I die", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "%s", json)
}

func main() {
	h := new(StatusHandler)
	http.Handle("/status", h)
	http.ListenAndServe("localhost:4000", nil)
}
