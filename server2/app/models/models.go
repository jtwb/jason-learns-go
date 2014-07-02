package models

import (
	"io/ioutil"
	"os"
)

type Status struct {
	Status   string
	Message  string
	Hostname string
	Version  string
	Kernel   string
	Errors   []string
}

func NewStatus() *Status {
	s := new(Status)

	s.getHostname()
	s.getVersion()
	s.getKernel()

	switch len(s.Errors) {
	case 0:
		s.Status = "OK"
	default:
		s.Status = "ERROR"
	}

	return s
}

func (s *Status) getHostname() {
	hostname, err := os.Hostname()
	if err != nil {
		s.Errors = append(s.Errors, err.Error())
	}
	s.Hostname = hostname
}

func (s *Status) getVersion() {
	version, err := ioutil.ReadFile("/proc/sys/kernel/version")
	if err != nil {
		s.Errors = append(s.Errors, err.Error())
	}
	s.Version = string(version)
}

func (s *Status) getKernel() {
	kernel, err := ioutil.ReadFile("/proc/sys/kernel/osrelease")
	if err != nil {
		s.Errors = append(s.Errors, err.Error())
	}
	s.Kernel = string(kernel)
}
