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
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/locales/zh_CN"
	"github.com/energye/lcl/logger"
	"github.com/energye/lcl/pkgs/i18n"
	"github.com/energye/lcl/pkgs/macapp"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/rtl/version"
	"github.com/energye/lcl/tools"
	"github.com/energye/lcl/tools/exec"
	"os"
	"path"
)

const (
	emfsLibsPath = "libs"
)

// Init LCL Global initialization
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	emfs.SetEMFS(libs, resources)
	if libname.LibName == "" {
		libname.LibName = libPath()
		if libname.LibName == "" {
			libname.LibName = path.Join(exec.HomeGoLCLDir, libname.GetDLLName())
			//lib If none of them exist, try to retrieve them from the built-in libs and release them to the user directory
			releaseLib(path.Join(emfsLibsPath, libname.GetDLLName()), libname.LibName)
			if !tools.IsExist(libname.LibName) {
				println(`Hint:
	Dependency library liblcl was not found
	If local liblcl exist, please put it in the specified location, If it does not exist, please download it from the Energy official website.
	Configuration Location:
		1. Current program execution directory
		2. USER_HOME/golcl/
		3. Environment variables LCL_HOME or ENERGY_HOME
			environment variable LCL_HOME is configured preferentially in the non-energy framework
			environment variable ENERGY_HOME takes precedence in the Energy framework
			ENERGY_HOME environment variable is recommended
`)
			}
		}
	}
	logger.Debug("LCL Init Lib Path:", libname.LibName)
	InitAll()
}

// If the lib dynamic library is built into EXE, release lib to the out directory in EXE
func releaseLib(fsPath, out string) {
	if emfs.GetLibsFS() != nil {
		// Attempt to create a directory, if the directory already exists, it will not be created
		tools.MkdirAll(exec.HomeGoLCLDir)
		var liblcl, err = emfs.GetLibsFS().ReadFile(fsPath)
		if err == nil {
			var file *os.File
			file, err = os.OpenFile(out, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			file.Write(liblcl)
		}
	}
}

func InitAll() {
	//macapp
	macapp.MacApp.Init()
	//api
	api.APIInit()
	//rtl
	rtl.RtlInit()
	//version
	version.VersionInit()
	//zh_cn
	zh_CN.ZH_CNInit()
	//win
	winInit()
	//lcl
	lcl.LCLInit()
	//i18n
	i18n.I18NInit()

}
