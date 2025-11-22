//----------------------------------------
//
// Copyright © yanghy. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package lcl

import (
	"github.com/energye/lcl/api"
	"github.com/energye/lcl/api/imports"
	"github.com/energye/lcl/base"
	"github.com/energye/lcl/types"
)

// ITrayIcon Parent: ICustomTrayIcon
type ITrayIcon interface {
	ICustomTrayIcon
}

type TTrayIcon struct {
	TCustomTrayIcon
}

// NewTrayIcon class constructor
func NewTrayIcon(theOwner IComponent) ITrayIcon {
	r := trayIconAPI().SysCallN(0, base.GetObjectUintptr(theOwner))
	return AsTrayIcon(r)
}

func TTrayIconClass() types.TClass {
	r := trayIconAPI().SysCallN(1)
	return types.TClass(r)
}

var (
	trayIconOnce   base.Once
	trayIconImport *imports.Imports = nil
)

func trayIconAPI() *imports.Imports {
	trayIconOnce.Do(func() {
		trayIconImport = api.NewDefaultImports()
		trayIconImport.Table = []*imports.Table{
			/* 0 */ imports.NewTable("TTrayIcon_Create", 0), // constructor NewTrayIcon
			/* 1 */ imports.NewTable("TTrayIcon_TClass", 0), // function TTrayIconClass
		}
	})
	return trayIconImport
}
