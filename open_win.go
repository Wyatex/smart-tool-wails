//go:build windows
// +build windows

package main

import (
	"os/exec"
	"strings"
	"syscall"
)

// open some uri
func (a *App) Open(value string) string {
	args := strings.Split(value, " ")
	if len(args) > 1 {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		err := cmd.Run()
		if err != nil {
			return err.Error()
		}
		return ""
	}
	cmd := exec.Command("cmd", "/C", "start", args[0])
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := cmd.Run()
	if err != nil {
		return err.Error()
	}
	return ""
}
