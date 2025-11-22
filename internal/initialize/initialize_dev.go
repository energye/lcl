// ----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// # Licensed under Apache License Version 2.0, January 2004
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// ----------------------------------------

//go:build !prod
// +build !prod

package initialize

import (
	"os"
	"path"
	"path/filepath"

	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/config"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/internal/initialize/macapp"
	"github.com/energye/lcl/tool"
	"github.com/energye/lcl/tool/exec"
)

// 开发环境加载 lib，在程序启动后读取用户目录/.energy 配置文件
// 不同操作系统加载方式也不同
// 优先级: 1. 自定义, 2. 当前执行目录, 3. .energy 配置文件 4. 相对目录
// 完整目录: [root]+energy+[framework]
func loadLibLCL(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	if tool.IsDarwin() {
		// 开发模式 自动生成 xxx.app
		macapp.Init()
	}
	// LCL 初始化时回调， 返回 lib 地址
	api.SetOnLoadLibCallback(func() (lib imports.DLL, err error) {
		libPath := libname.LibName
		if libPath != "" {
			// 自定义加载目录
			lib, err = imports.NewDLL(libPath)
		} else if tool.IsDarwin() {
			// MacOS 固定加载目录
			libPath = "@executable_path/../Frameworks/" + libname.GetDLLName()
		} else {
			// Windows, Linux
			// 优先当前执行目录
			currentPathLibPath := path.Join(exec.Dir, libname.GetDLLName())
			if tool.IsExist(currentPathLibPath) {
				libPath = currentPathLibPath
			} else {
				// 开发环境配置目录
				if config.Get() != nil {
					libPath = filepath.Join(config.Get().FrameworkPath(), libname.GetDLLName())
				} else {
					// 最后尝试相对目录
					libPath = libname.GetDLLName()
				}
			}
			if libPath == "" {
				// 最后 尝试从临时目录获取
				libPath = filepath.Join(os.TempDir(), libname.GetDLLName())
			}
		}
		// 加载 LibLCL
		if libPath != "" {
			libname.LibName = libPath
			lib, err = imports.NewDLL(libPath)
		}
		if lib == 0 {
			if err != nil {
				println("[ERROR] Load LibLCL", err.Error())
			}
			println("[Hint] LibLCL Path:", libname.LibName)
			panic(`[Hint]:
  Failed initialize LibLCL, check the development environment
  Use CLI: 
    [energy env] : Check the configuration of the development environment
    [energy install] : Installation development environment
`)
		}
		return
	})
	// 加载 lib dll
	api.Init()
}
