//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	. "github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	. "github.com/energye/lcl/types"
)

// ITrayIcon Parent: ICustomTrayIcon
type ITrayIcon interface {
	ICustomTrayIcon
}

// TTrayIcon Parent: TCustomTrayIcon
type TTrayIcon struct {
	TCustomTrayIcon
}

func NewTrayIcon(TheOwner IComponent) ITrayIcon {
	r1 := rayIconImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsTrayIcon(r1)
}

func TrayIconClass() TClass {
	ret := rayIconImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	rayIconImport       *imports.Imports = nil
	rayIconImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("TrayIcon_Class", 0),
		/*1*/ imports.NewTable("TrayIcon_Create", 0),
	}
)

func rayIconImportAPI() *imports.Imports {
	if rayIconImport == nil {
		rayIconImport = NewDefaultImports()
		rayIconImport.SetImportTable(rayIconImportTables)
		rayIconImportTables = nil
	}
	return rayIconImport
}
