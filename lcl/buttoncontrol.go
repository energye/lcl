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

// IButtonControl Parent: IWinControl
type IButtonControl interface {
	IWinControl
}

// TButtonControl Parent: TWinControl
type TButtonControl struct {
	TWinControl
}

func NewButtonControl(TheOwner IComponent) IButtonControl {
	r1 := buttonControlImportAPI().SysCallN(1, GetObjectUintptr(TheOwner))
	return AsButtonControl(r1)
}

func ButtonControlClass() TClass {
	ret := buttonControlImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	buttonControlImport       *imports.Imports = nil
	buttonControlImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("ButtonControl_Class", 0),
		/*1*/ imports.NewTable("ButtonControl_Create", 0),
	}
)

func buttonControlImportAPI() *imports.Imports {
	if buttonControlImport == nil {
		buttonControlImport = NewDefaultImports()
		buttonControlImport.SetImportTable(buttonControlImportTables)
		buttonControlImportTables = nil
	}
	return buttonControlImport
}
