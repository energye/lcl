//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package imports

import (
	"syscall"
)

func NewDLL(name string) (DLL, error) {
	h, err := syscall.LoadLibrary(name)
	return DLL(h), err
}

func (h DLL) Release() error {
	if h != 0 {
		return syscall.FreeLibrary(syscall.Handle(h))
	}
	return nil
}

func (h DLL) Proc(name string) (ProcAddr, error) {
	addr, err := syscall.GetProcAddress(syscall.Handle(h), name)
	return ProcAddr(addr), err
}

func (p ProcAddr) Call(args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return SyscallN(uintptr(p), args...)
}

func SyscallN(addr uintptr, args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
	if addr == 0 {
		return 0, 0, syscall.ERROR_PROC_NOT_FOUND
	}
	return syscall.SyscallN(addr, args...)
}
