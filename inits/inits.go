//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package inits

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/inits/config"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/locales/zh_CN"
	"github.com/energye/lcl/pkgs/i18n"
	"github.com/energye/lcl/pkgs/macapp"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/rtl/version"
	"github.com/energye/lcl/tools"
	"github.com/energye/lcl/tools/exec"
	"path"
	"path/filepath"
)

// Init LCL Global initialization
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	// MacOS 开发模式 自动生成 xxx.app
	macapp.Init()
	api.SetOnLoadLibCallback(func() (dll imports.DLL, err error) {
		libdllPath := libname.LibName
		if libdllPath != "" {
			// 自定义加载目录
			dll, err = imports.NewDLL(libdllPath)
		} else if tools.IsDarwin() {
			// MacOS 固定加载目录
			libdllPath = "@executable_path/../Frameworks/" + libname.GetDLLName()
		} else {
			// Windows, Linux
			// 优先当前执行目录
			currentPathLibPath := path.Join(exec.Dir, libname.GetDLLName())
			if tools.IsExist(currentPathLibPath) {
				libdllPath = currentPathLibPath
			} else {
				// 开发环境配置目录
				if config.Get() != nil {
					libdllPath = filepath.Join(config.Get().FrameworkPath(), libname.GetDLLName())
				} else {
					// 最后尝试相对目录
					libdllPath = libname.GetDLLName()
				}
			}
		}
		// 加载 LibLCL
		if libdllPath != "" {
			libname.LibName = libdllPath
			dll, err = imports.NewDLL(libdllPath)
		}
		if dll == 0 {
			if err != nil {
				println("Load LibLCL Error:", err.Error())
			}
			println("LibLCL Path:", libname.LibName)
			panic(`Hint:
  Failed initialize LibLCL, check the development environment
  Use CLI: 
    [energy env] : Check the configuration of the development environment
    [energy install] : Installation development environment
`)
		}
		return
	})
	emfs.SetEMFS(libs, resources)
	InitAll()
}

func InitAll() {
	// api
	api.APIInit()
	// rtl
	rtl.RtlInit()
	// version
	version.VersionInit()
	// zh_cn
	zh_CN.ZH_CNInit()
	// win
	APIInit()
	// lcl
	lcl.LCLInit()
	// i18n
	i18n.I18NInit()

}
