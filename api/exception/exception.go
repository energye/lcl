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
//	Supports: Windows, MacOS.
package exception

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/internal/exception"
)

// SetOnException
// 捕获底层库异常
func SetOnException(fn exception.Callback) {
	if exception.HandlerCallback == nil {
		api.SetExceptionHandlerCallback(exceptionHandlerProcEventAddr)
		exception.HandlerCallback = fn
	}
}

func exceptionHandlerProc(funcName, message uintptr) uintptr {
	if exception.HandlerCallback != nil {
		exception.HandlerCallback(api.GoStr(funcName), api.GoStr(message))
	}
	return 0
}
