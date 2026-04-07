//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// :predefine:

//go:build !windows && !cgo

package lcl

import "github.com/ebitengine/purego"

var (
	eventCallback        = purego.NewCallback(eventCallbackProc)
	removeEventCallback  = purego.NewCallback(removeEventCallbackProc)
	messageCallback      = purego.NewCallback(messageCallbackProc)
	threadSyncCallback   = purego.NewCallback(threadSyncCallbackProc)
	threadAsyncCallback  = purego.NewCallback(threadAsyncCallbackProc)
	createParamsCallback = purego.NewCallback(createParamsCallbackProc)
	formCreateCallback   = purego.NewCallback(formCreateCallbackProc)
)
