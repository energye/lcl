//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//	Underlying dynamic link library exception capture

package api

import (
	"strconv"

	"github.com/energye/lcl/api/internal/exception"
)

var (
	registerExceptionHandler bool
)

func SetDebug(value bool) {
	exception.Debug = value
}

func Debug() bool {
	return exception.Debug
}

// SetOnException	`
// Base library exception callback
func SetOnException(fn exception.Callback) {
	exception.SetOnException(fn)
	if !registerExceptionHandler {
		registerExceptionHandler = true
		SetEventCallback(exceptionHandlerProcEventAddr, EctExceptionHandler)
	}
}

func exceptionHandlerProc(idType uintptr, message uintptr) uintptr {
	exception.CallOnException(strconv.Itoa(int(idType)), GoStr(message))
	return 0
}
