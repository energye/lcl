//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
//----------------------------------------

//go:build prod

package initialize

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/tool/exec"
	"os"
	"path/filepath"
)

// 发布环境加载 libENERGY，不再依赖 .energy 配置文件
// 优先级: 1. 自定义, 2. 当前执行目录, 3. 相对目录
// 不同操作系统加载方式也不同
func loadLibENERGY(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	// LCL 初始化时回调， 返回 lib 地址
	api.SetOnLoadLibCallback(func() (lib imports.DLL, err error) {
		libPath := libname.LibName
		if libPath != "" {
			// 自定义加载目录
			// lib, err = imports.NewDLL(libPath)
		} else if tool.IsDarwin() {
			// MacOS 固定加载目录
			isUniversal := os.Getenv("--universal-binary") == "universal"
			if isUniversal {
				libPath = "@executable_path/../Frameworks/libenergy-darwin-universal-cocoa.dylib"
			} else {
				libPath = "@executable_path/../Frameworks/" + libname.GetDLLName()
			}
		} else {
			// Windows, Linux
			if currentPathLibPath := filepath.Join(exec.Dir, libname.GetDLLName()); tool.IsExist(currentPathLibPath) {
				// 优先当前执行目录
				libPath = currentPathLibPath
			} else {
				// 最后尝试相对目录
				libPath = libname.GetDLLName()
			}
		}
		// 加载 LibENERGY
		if libPath != "" {
			libname.LibName = libPath
			lib, err = imports.NewDLL(libPath)
		}
		if lib == 0 {
			if err != nil {
				println("[ERROR] Load LibENERGY", err.Error())
			}
			println("[ERROR] Path:", libname.LibName)
			panic(`Failed initialize LibENERGY`)
			os.Exit(1)
		}
		return
	})
	// 加载 lib dll
	api.Init()
}
