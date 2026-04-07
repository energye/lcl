//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows && !cgo

package imports

import (
	"github.com/ebitengine/purego"
)

func NewDLL(name string) (DLL, error) {
	r, err := purego.Dlopen(name, purego.RTLD_GLOBAL|purego.RTLD_LAZY)
	return DLL(r), err
}

func (h DLL) Release() error {
	return purego.Dlclose(uintptr(h))
}

func (h DLL) Proc(name string) (ProcAddr, error) {
	addr, err := purego.Dlsym(uintptr(h), name)
	return ProcAddr(addr), err
}

func (p ProcAddr) Call(args ...uintptr) (r1, r2 uintptr, err error) {
	return SyscallN(uintptr(p), args...)
}

func SyscallN(addr uintptr, args ...uintptr) (r1, r2 uintptr, err error) {
	//var r3 uintptr
	r1, r2, _ = purego.SyscallN(addr, args...)
	return
}

func NewCallback(fn any) uintptr {
	return purego.NewCallback(fn)
}
