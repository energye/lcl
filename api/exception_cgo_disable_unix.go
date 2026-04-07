//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build !windows && !cgo

package api

import "github.com/energye/lcl/api/imports"

var (
	exceptionHandlerProcEventAddr = imports.NewCallback(exceptionHandlerProc)
)
