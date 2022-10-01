//go:build linux
// +build linux

package main

import (
	"os/exec"
	"strings"
)

// open some uri
func (a *App) Open(value string) string {
	args := strings.Split(value, " ")
	if len(args) > 1 {
		cmd := exec.Command(args[0], args[1:]...)
		err := cmd.Run()
		if err != nil {
			return err.Error()
		}
		return ""
	}
	cmd := exec.Command("xdg-open", args[0])
	err := cmd.Run()
	if err != nil {
		return err.Error()
	}
	return ""
}
