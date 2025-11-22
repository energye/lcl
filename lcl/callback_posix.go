//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build !windows && cgo
// +build !windows,cgo

package lcl

//// #cgo darwin CFLAGS: -mmacosx-version-min=10.15 -DMACOSX_DEPLOYMENT_TARGET=10.15
// #cgo darwin CFLAGS: -mmacosx-version-min=10.15
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.15
// #include <stdint.h>
//
// extern void* doEventCallbackProc(uintptr_t f, void* args, long argcount);
// static void* doGetEventCallbackAddr() {
//    return &doEventCallbackProc;
// }
//
// extern void* doRemoveEventCallbackProc(uintptr_t ptr);
// static void* doRemoveEventCallbackAddr() {
//    return &doRemoveEventCallbackProc;
// }
//
// extern void* doMessageCallbackProc(uintptr_t f, void* msg);
// static void* doMessageCallbackAddr() {
//    return &doMessageCallbackProc;
// }
//
// extern void* doThreadSyncCallbackProc();
// static void* doThreadSyncCallbackAddr() {
//    return &doThreadSyncCallbackProc;
// }
//
// extern void* doThreadAsyncCallbackProc(void* data);
// static void* doThreadAsyncCallbackAddr() {
//    return &doThreadAsyncCallbackProc;
// }
//
// extern void* doCreateParamsCallbackProc(uintptr_t ptr, void* sender, void* params);
// static void* doCreateParamsCallbackAddr() {
//    return &doCreateParamsCallbackProc;
// }
//
// extern void* doFormCreateCallbackProc(uintptr_t ptr, void* sender);
// static void* doFormCreateCallbackAddr() {
//    return &doFormCreateCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doEventCallbackProc
func doEventCallbackProc(f C.uintptr_t, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(uintptr(f), uintptr(args), int(argcount))
	return nil
}

//export doRemoveEventCallbackProc
func doRemoveEventCallbackProc(ptr C.uintptr_t) unsafe.Pointer {
	removeEventCallbackProc(uintptr(ptr))
	return nil
}

//export doMessageCallbackProc
func doMessageCallbackProc(f C.uintptr_t, msg unsafe.Pointer) unsafe.Pointer {
	messageCallbackProc(uintptr(f), uintptr(msg))
	return nil
}

//export doThreadSyncCallbackProc
func doThreadSyncCallbackProc() unsafe.Pointer {
	threadSyncCallbackProc()
	return nil
}

//export doThreadAsyncCallbackProc
func doThreadAsyncCallbackProc(data unsafe.Pointer) unsafe.Pointer {
	threadAsyncCallbackProc(uintptr(data))
	return nil
}

//export doCreateParamsCallbackProc
func doCreateParamsCallbackProc(ptr C.uintptr_t, sender, params unsafe.Pointer) unsafe.Pointer {
	createParamsCallbackProc(uintptr(ptr), uintptr(sender), uintptr(params))
	return nil
}

//export doFormCreateCallbackProc
func doFormCreateCallbackProc(ptr C.uintptr_t, sender unsafe.Pointer) unsafe.Pointer {
	formCreateCallbackProc(uintptr(ptr), uintptr(sender))
	return nil
}

var (
	eventCallback        = uintptr(C.doGetEventCallbackAddr())
	removeEventCallback  = uintptr(C.doRemoveEventCallbackAddr())
	messageCallback      = uintptr(C.doMessageCallbackAddr())
	threadSyncCallback   = uintptr(C.doThreadSyncCallbackAddr())
	threadAsyncCallback  = uintptr(C.doThreadAsyncCallbackAddr())
	createParamsCallback = uintptr(C.doCreateParamsCallbackAddr())
	formCreateCallback   = uintptr(C.doFormCreateCallbackAddr())
)
