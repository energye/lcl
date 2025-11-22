//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

package lcl

import (
	"syscall"
)

var (
	eventCallback        = syscall.NewCallback(eventCallbackProc)
	removeEventCallback  = syscall.NewCallback(removeEventCallbackProc)
	messageCallback      = syscall.NewCallback(messageCallbackProc)
	threadSyncCallback   = syscall.NewCallback(threadSyncCallbackProc)
	threadAsyncCallback  = syscall.NewCallback(threadAsyncCallbackProc)
	createParamsCallback = syscall.NewCallback(createParamsCallbackProc)
	formCreateCallback   = syscall.NewCallback(formCreateCallbackProc)
)
