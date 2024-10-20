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

// ICustomUpDown Parent: ICustomControl
type ICustomUpDown interface {
	ICustomControl
}

// TCustomUpDown Parent: TCustomControl
type TCustomUpDown struct {
	TCustomControl
}

func NewCustomUpDown(AOwner IComponent) ICustomUpDown {
	r1 := customUpDownImportAPI().SysCallN(1, GetObjectUintptr(AOwner))
	return AsCustomUpDown(r1)
}

func CustomUpDownClass() TClass {
	ret := customUpDownImportAPI().SysCallN(0)
	return TClass(ret)
}

var (
	customUpDownImport       *imports.Imports = nil
	customUpDownImportTables                  = []*imports.Table{
		/*0*/ imports.NewTable("CustomUpDown_Class", 0),
		/*1*/ imports.NewTable("CustomUpDown_Create", 0),
	}
)

func customUpDownImportAPI() *imports.Imports {
	if customUpDownImport == nil {
		customUpDownImport = NewDefaultImports()
		customUpDownImport.SetImportTable(customUpDownImportTables)
		customUpDownImportTables = nil
	}
	return customUpDownImport
}
