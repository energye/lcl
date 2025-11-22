//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

// Package exception
//
//	Underlying dynamic link library exception capture
package exception

import (
	"github.com/energye/lcl/api"
)

type Callback func(exception int32, message string)

var exceptionCallback Callback

// SetOnException	`
// Base library exception callback
func SetOnException(fn Callback) {
	if exceptionCallback == nil {
		api.SetEventCallback(exceptionHandlerProcEventAddr, api.EctExceptionHandler)
		exceptionCallback = fn
	}
}

func exceptionHandlerProc(idType uintptr, message uintptr) uintptr {
	if exceptionCallback != nil {
		exceptionCallback(int32(idType), api.GoStr(message))
	}
	return 0
}
