//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows

package command

import (
	"syscall"
)

func HideWindow(bool bool) *syscall.SysProcAttr {
	return &syscall.SysProcAttr{HideWindow: bool}
}

func (m *CMD) Kill() error {
	return m.Cmd.Process.Kill()
	//return exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(m.Cmd.Process.Pid)).Run()
}
