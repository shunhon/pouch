package modules

import (
	"os"
	"strings"
)

type Cwd struct {
	Module
}

func (m *Cwd) Init() {
	m.color = "blue"
}

func (m *Cwd) GetOutput() string {
	// TODO: Fix error handling
	cwd, _ := os.Getwd()
	home := os.Getenv("HOME")
	if strings.HasPrefix(cwd, home) {
		cwd = strings.Replace(cwd, home, "~", 1)
	}
	return cwd
}
