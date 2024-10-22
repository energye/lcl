//go:build darwin
// +build darwin

//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package inits

import (
	"github.com/energye/lcl/api/libname"
)

func winInit() {
}

func libPath() string {
	return "@executable_path/../Frameworks/" + libname.GetDLLName()
}
