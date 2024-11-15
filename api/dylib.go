//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/api/libname"
)

var (
	uiLib           imports.DLL                 // 全局导入库
	loadLibCallback func() (imports.DLL, error) // 自定义加载liblcl动态库回调函数
)

// SetOnLoadLibCallback
//
//	加载liblcl动态库回调函数
//	调用 inits.Init 之前设置
func SetOnLoadLibCallback(fn func() (imports.DLL, error)) {
	if loadLibCallback == nil {
		loadLibCallback = fn
	}
}

func loadUILib() imports.DLL {
	if loadLibCallback != nil {
		lib, err := loadLibCallback()
		if lib == 0 && err != nil {
			panic(err)
		}
		return lib
	} else {
		lib, err := imports.NewDLL(libname.LibName)
		if lib == 0 && err != nil {
			panic(err)
		}
		return lib
	}
}

func closeLib() {
	if uiLib != 0 {
		uiLib.Release()
		uiLib = 0
	}
}

// 调用预定义导入API
func defSyscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := LCLPreDef().Proc(trap).Call(args...)
	return r1
}
