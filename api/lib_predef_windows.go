//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build windows
// +build windows

package api

import (
	"github.com/energye/lcl/api/imports"
	"sync"
)

func CreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	platformPreDefAPI().SysCallN(0, PasStr(aDestPath), PasStr(aShortCutName), PasStr(aURL))
}

func CreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	return GoBool(platformPreDefAPI().SysCallN(1, PasStr(aDestPath), PasStr(aShortCutName), PasStr(aSrcFileName),
		PasStr(aIconFileName), PasStr(aDescription), PasStr(aCmdArgs)))
}

var (
	libPlatformPreDefOnce   sync.Once
	libPlatformPreDefImport *imports.Imports
)

func platformPreDefAPI() *imports.Imports {
	libPlatformPreDefOnce.Do(func() {
		libPlatformPreDefImport = NewDefaultImports()
		libPlatformPreDefImport.Table = []*imports.Table{
			imports.NewTable("DCreateURLShortCut", 0),
			imports.NewTable("DCreateShortCut", 0),
		}
	})
	return libPlatformPreDefImport
}
