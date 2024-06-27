//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows && cgo
// +build !windows,cgo

package exception

//// #cgo darwin CFLAGS: -mmacosx-version-min=10.8 -DMACOSX_DEPLOYMENT_TARGET=10.8
// #cgo darwin CFLAGS: -mmacosx-version-min=10.10
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.10
//
// extern void* doExceptionHandlerProc(void* funcName, void* message);
// static void* doExceptionHandlerProcEventAddr() {
//    return &doExceptionHandlerProc;
// }
import "C"

import (
	"unsafe"
)

//export doExceptionHandlerProc
func doExceptionHandlerProc(funcName, message unsafe.Pointer) unsafe.Pointer {
	exceptionHandlerProc(uintptr(funcName), uintptr(message))
	return nil
}

var (
	exceptionHandlerProcEventAddr = uintptr(C.doExceptionHandlerProcEventAddr())
)
