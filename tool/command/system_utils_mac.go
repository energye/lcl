//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build darwin

package command

import "syscall"

func HideWindow(bool bool) *syscall.SysProcAttr {
	return nil
}

func (m *CMD) Kill() error {
	return m.Cmd.Process.Kill()
}
