//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package inits

import (
	"errors"
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/api/libname"
	"github.com/energye/lcl/emfs"
	"github.com/energye/lcl/lcl"
	"github.com/energye/lcl/locales/zh_CN"
	"github.com/energye/lcl/logger"
	"github.com/energye/lcl/pkgs/i18n"
	"github.com/energye/lcl/pkgs/macapp"
	"github.com/energye/lcl/rtl"
	"github.com/energye/lcl/rtl/version"
	"github.com/energye/lcl/tools/exec"
	"os"
	"path"
)

// Init LCL Global initialization
func Init(libs emfs.IEmbedFS, resources emfs.IEmbedFS) {
	emfs.SetEMFS(libs, resources)
	api.SetOnLoadLibCallback(func() (dll imports.DLL, err error) {
		// current user home dir load
		if dllPath, ok := releaseLib(path.Join("libs", libname.GetDLLName()), path.Join(exec.HomeDir, libname.GetDLLName())); ok {
			dll, err = imports.NewDLL(dllPath)
			logger.Debug("LCL Init Lib Path:", dllPath)
			libname.LibName = ""
			return
		} else if dllPath = libPath(); dllPath != "" {
			// default load
			dll, err = imports.NewDLL(dllPath)
			logger.Debug("LCL Init Lib Path:", dllPath)
			libname.LibName = ""
			return
		}
		return 0, errors.New(`Hint:
	Failed to load liblcl, liblcl not found.
	Please configure the liblcl environment correctly.
	Load Priority: 
		level 1: Embedded into exe
		level 2: libname.LibName = "dll full path"
		level 3: Current_Directory > Environment_Variable(LCL_HOME | LCLCEF_HOME | LCLWV2_HOME | LCLWK2_HOME | ENERGY_HOME) > $USER_HOME/.energy Read from file
`)
	})
	InitAll()
}

// If the lib dynamic library is built into EXE, release lib to the out directory in EXE
func releaseLib(fsPath, outDllPath string) (string, bool) {
	if emfs.GetLibsFS() != nil {
		// Attempt to create a directory, if the directory already exists, it will not be created
		var dllData, err = emfs.GetLibsFS().ReadFile(fsPath)
		if err == nil {
			var file *os.File
			file, err = os.OpenFile(outDllPath, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()
			file.Write(dllData)
			return outDllPath, true
		}
	}
	return "", false
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
